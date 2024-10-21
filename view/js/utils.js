const Msg_Type = Object.freeze({
    event: "EVENT",
    ok: "OK",
    eose: "EOSE",
    closed: "CLOSED",
    notice:  "NOTICE",
});

function ele_get(query) {
    return document.querySelector(query);
}
function eles_get(query) {
    return document.querySelectorAll(query);
}

function hexToBytes(hex) {
    return Uint8Array.from(hex.match(/.{1,2}/g).map((byte) => parseInt(byte, 16)));
}
function bytes_to_hex(bytes) {
    return bytes.reduce( ( str, byte ) => str + byte.toString( 16 ).padStart( 2, "0" ), "" );
}

function date_now() {
    return Math.floor(Date.now() / 1000);
}
function date_format(timestamp) {
    const date = new Date(timestamp * 1000); // Convert to milliseconds

    const options = {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: true, // Use 12-hour format
    };

    return date.toLocaleString('en-GB', options).replace(',', ' ·').replaceAll(" ", "");
}

function gen_prikey() {
    return bytes_to_hex(nobleSecp256k1.utils.randomPrivateKey());
}
function get_prikey() {
    return localStorage.getItem("prikey");
}
function set_prikey(key) {
    localStorage.setItem("prikey", key);
}
function del_prikey() {
    localStorage.removeItem("prikey");
}

function get_pubkey() {
    const prikey = get_prikey();
    return nobleSecp256k1.getPublicKey(prikey, true).substring(2);
}

function set_content(content) {
    localStorage.setItem("content", content)
}
function get_content() {
    return localStorage.getItem("content");
}
function del_content() {
    return localStorage.removeItem("content");
}

function gen_subid() {
    return bytes_to_hex(nobleSecp256k1.utils.randomPrivateKey()).substring(0, 16);
}
async function gen_eventid(data) {
    return bytes_to_hex(await sha256((new TextEncoder().encode(data))));
}


function is_login() {
    return get_prikey() ? true : false;
}
