import {RelayPool} from 'https://esm.sh/nostr@0.2.8';

const pubkey = "32e1827635450ebb3c5a7d12c1f8e7b2b514439ac10a67eef3d9fd9c5c68e245"
const damus = "wss://relay.damus.io"
const scsi = "wss://nostr-pub.wellorder.net"
const relays = [damus, scsi]

const pool = RelayPool(relays)

pool.on('open', relay => {
    if (get_prikey()) {
        const filter  = { "authors": [ pubkey ], "kinds": [ 1 ], "limit": 10 };
        // const filter  = { "kinds": [ 1 ], "limit": 1 };
        const id = gen_subid();
        relay.subscribe(id, filter);
        setup_login_profile();
    } else {
        const filter  = { "kinds": [ 1 ], "limit": 1 };
        const id = gen_subid();
        relay.subscribe(id, filter);
    }
});

pool.on('event', (relay, sub_id, event) => {
    if (event.kind == 0){
        const { id, pubkey, content } = event;
        const {
            name, display_name, about, picture, website, banner, bot
        } = JSON.parse(content);
        set_profile(
            pubkey, name, display_name, about, picture, website, banner, bot
        );
    }
    if (event.kind == 1) {
        const { content, created_at, id, pubkey, tags } = event;

        const e_tag = tags[0]
        if (e_tag.includes("e")) {
            console.log(e_tag[1]);
        } else {
            createPost(content, created_at, id, pubkey);
            const filter_info = { "kinds": [ 0 ], authors: [ pubkey ] };
            send_req(filter_info);
            const filter_reaction = { "kinds": [ 1, 7 ], "#e": [ id ] };
            send_req(filter_reaction);
        }
    }
    if (event.kind == 7) {
        const tag = event.tags[0];
        if (tag.includes("e")) {
            const post = document.querySelector(`.post_${tag[1]}`);
            console.log(event);
            if (post) {
                const post_react_count = post.querySelector(".actions_react_count");
                const current_count = post_react_count.innerHTML;
                post_react_count.innerHTML = Number(current_count) + 1;
            }
        }
    }
});

pool.on('eose', (relay, id) => {
    // relay.close()
    console.log("EOSE:", relay.url, id);
});

pool.on('notice', (relay, msg) => {
    console.log("Notice:", msg);
});

pool.on('ok', (relay, msg) => {
    console.log("Notice:", msg);
});

pool.on('closed', (relay, msg) => {
    console.log("Closed:", msg);
});

function send_req(filter) {
    const id = gen_subid();
    const subscription = [ "REQ", id, filter ];
    pool.send(subscription);
}

function setup_login_profile() {
    const pubkey = get_pubkey();
    ele_get("#account .account_main_picture")
        .className += ` profile_picture_${pubkey}`;
    ele_get("#account .account_main_display_name")
        .className += ` profile_display_name_${pubkey}`;
    ele_get("#account .account_main_username")
        .className += ` profile_username_${pubkey}`;
    ele_get("#account .account_main_about")
        .className += ` profile_about_${pubkey}`;

    ele_get("#account_edit_display_name")
        .className += ` profile_username_${pubkey}`;
    ele_get("#account_edit_status")
        .className += ` profile_status_${pubkey}`;
    ele_get("#account_edit_banner")
        .className += ` profile_banner_${pubkey}`;
    ele_get("#account_edit_picture")
        .className += ` profile_picture_${pubkey}`;
    ele_get("#account_edit_about")
        .className += ` profile_about_${pubkey}`;

    // ele_get("#account .account_main_username").textContent = pubkey;
    send_req({ "kinds": [ 0 ], "limit": 1, authors: [ pubkey ] });
}

function edit_profile() {
}
window.edit_profile = edit_profile;

async function send_note() {
    const content = ele_get("#new_post_textbox").value;
    const sig_note = await get_sig_note(content);
    ele_get("#new_post_textbox").value = "";
    del_content();
    const event = ["EVENT", sig_note];
    pool.send(event);
}
window.send_note = send_note;

async function send_reaction(emoji, id, send_pubkey) {
    const pubkey = get_pubkey();

    const setup_event = {
        "pubkey": pubkey,
        "created_at": date_now(),
        "kind": 7,
        "tags": [
            [ "e", id ],
            [ "p", send_pubkey ],
        ],
        "content": emoji
    }

    const sig_reaction = await get_sig_event(setup_event);
    const event = ["EVENT", sig_reaction];
    pool.send(event);
}
window.send_reaction = send_reaction;
