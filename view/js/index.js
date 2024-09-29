const relayUrls = [
    'wss://nos.lol',
    'wss://relay.nostr.net'
];

const sockets = [];

function connectToRelay(url) {

const ws = new WebSocket(url);

ws.addEventListener("open", async (e) => {
    console.log("connected to " + url);

    if (get_prikey()) {
        const filter  = { "authors": [ get_pubkey() ], "kinds": [ 1 ], "limit": 10 };
        // const filter  = { "kinds": [ 1 ], "limit": 1 };

const id = gen_subid();
const subscription = [ "REQ", id, filter ];
ws.send(JSON.stringify(subscription));
        // send_req(filter);

        // setup_login_profile();
    } else {
        const filter  = { "kinds": [ 1 ], "limit": 1 };

const id = gen_subid();
const subscription = [ "REQ", id, filter ];
ws.send(JSON.stringify(subscription));
        // send_req(filter);
    }
});

ws.addEventListener("message", async (msg) => {
    const data = JSON.parse(msg.data);
    const msg_type = data[0];

    switch (msg_type) {
        case Msg_Type.event:
            const [ type, id, event ] = data;

            if (event.kind == 0){
                const { id, pubkey, content } = event;
                const { name, display_name, about, picture, website, banner, bot } = JSON.parse(content);
                set_profile(pubkey, name, display_name, about, picture, website, banner, bot);
            }
            if (event.kind == 1) {
                const { content, created_at, id, pubkey } = event;
                createPost(content, created_at, id, pubkey);
                const filter = { "kinds": [ 0 ], "limit": 1, authors: [ pubkey ] };


const subid = gen_subid();
const subscription = [ "REQ", subid, filter ];
ws.send(JSON.stringify(subscription));
                // send_req(filter);
            }

            break;
        case Msg_Type.ok:
            console.log("OK", data);
            break;
        case Msg_Type.eose:
            console.log("EOSE", data);
            break;
        case Msg_Type.closed:
            console.log("CLOSED", data);
            break;
        case Msg_Type.notice:
            console.log("NOTICE", data);
            break;
        default:
            console.log("Error: Message type is not defined.");
    }
});

ws.addEventListener("error", async (msg) => {
    console.log(msg);
});

ws.addEventListener("close", async (msg) => {
    console.log("Closed: " + msg);
});

return ws;
}

relayUrls.forEach(url => {
    const socket = connectToRelay(url);
    sockets.push(socket);
});

function send_req(filter) {
    const id = gen_subid();
    const subscription = [ "REQ", id, filter ];
sockets.forEach(ws => {
    if (ws.readyState === WebSocket.OPEN) ws.send(JSON.stringify(subscription));
});
}

async function send_note() {
    const content = ele_get("#new_post_textbox").value;
    const sig_note = await get_sig_note(content);
    ele_get("#new_post_textbox").value = "";
    del_content();

sockets.forEach(ws => {
    if (ws.readyState === WebSocket.OPEN) ws.send(JSON.stringify(["EVENT", sig_note]));
});
}

function login() {
    const prikey = ele_get("#pri_key").value;
    set_prikey(prikey);
    setup_login(true);
}

function logout() {
    del_prikey();
    setup_login(false);
}

function setup_login(opt) {
    if (opt) {
        ele_get("#login").style.display = "none";
        ele_get("#account").style.display = "";
        ele_get("#new_post").style.display = "";
    } else {
        ele_get("#login").style.display = "";
        ele_get("#account").style.display = "none";
        ele_get("#new_post").style.display = "none";
    }
}

function setup_login_profile() {
    const pubkey = get_pubkey();
    ele_get("#account .account_main_picture").className += ` profile_picture_${pubkey}`;
    ele_get("#account .account_main_display_name").className += ` profile_display_name_${pubkey}`;
    ele_get("#account .account_main_username").className += ` profile_username_${pubkey}`;
    ele_get("#account .account_main_about").className += ` profile_about_${pubkey}`;
    // ele_get("#account .account_main_username").textContent = pubkey;
    send_req({ "kinds": [ 0 ], "limit": 1, authors: [ pubkey ] });
}

function createPost(content, created_at, id, pubkey) {
    const article = document.createElement("article");
    article.className = `post_${id}`;
    content = DOMPurify.sanitize(marked.parse(content));
    const date = date_format(created_at);

    article.innerHTML =  `
<div class="post_profile">
    <div class="post_profile_image">
        <a><img class="profile_picture_${pubkey}" src="assets/default-profile.png" alt="Profile image ${pubkey}" /></a>
    </div>
    <div class="post_profile_info">
        <div class="post_profile_info_1">
            <a><span class="profile_display_name_${pubkey}"></span></a> <small><time class="time">${date}</time></small>
        </div>
        <div class="post_profile_info_2">
            <small class="profile_username_${pubkey}">${pubkey}</small>
        </div>
    </div>
</div>
<div class="post_content">
    ${content}
</div>
<div class="actions">
    <button class="actions_comment">Comment (3)</button>
    <button class="actions_react">React</button>
</div>`;

    ele_get("#posts").prepend(article);
}

function set_profile(pubkey, name, display_name, about, picture, website, banner, bot) {
    const fname = display_name || name
    eles_get(`.profile_picture_${pubkey}`).forEach(ele => {
        ele.src = picture || "assets/default-profile.png";
        ele.alt = `Profile image of ${fname}`;
    } );
    eles_get(`.profile_display_name_${pubkey}`).forEach(ele => ele.textContent = fname );
    eles_get(`.profile_about_${pubkey}`).forEach(ele => ele.textContent = about );
}

async function get_sig_note(content) {
    const pubkey = get_pubkey();

    const event = {
        "pubkey": pubkey,
        "created_at": date_now(),
        "kind": 1,
        "tags": [],
        "content": content,
    }

    return await get_sig_event(event);
}

async function get_sig_event(event) {
    const prikey = get_prikey();
    const event_data = JSON.stringify([
        0,                    // Reserved for future use
        event['pubkey'],      // The sender's public key
        event['created_at'],  // Unix timestamp
        event['kind'],        // Message “kind” or type
        event['tags'],        // Tags identify replies/recipients
        event['content']      // Your note contents
    ]);
    // event.id  = bytes_to_hex(await sha256((new TextEncoder().encode(eventData))));
    event.id  = await gen_eventid(event_data);
    event.sig = await schnorr.sign( event.id,  prikey);
    return event;
}

if (get_prikey()) {
    setup_login(true);
} else {
    not_login();
}
