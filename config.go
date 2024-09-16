package main

var config = map[string]interface{}{
    "name": "drift/social",
    "description": "Connect to people freely.",
    "pubkey": "cbc6fc60139278ad2919647117d2dabbd55c27cb84e93fd4728701c05e022c37",
    "contact": "https://AnzenKodo.github.io",
    "supported_nips": []string{"NIP-01", "NIP-11"},

    // "icon": "https://example.com/logo.png",
    // "tags": ["sfw-only", "bitcoin-only", "anime"],
    "language_tags": []string{"en", "en-419"},
    "relay_countries": []string{ "IN" },
    // "posting_policy": "https://example.com/posting-policy.html",
    // "limitation": {
    //     // "max_message_length": 16384,
    //     // "max_subscriptions": 20,
    //     // "max_filters": 100,
    //     // "max_limit": 5000,
    //     // "max_subid_length": 100,
    //     // "max_event_tags": 100,
    //     // "max_content_length": 8196,
    //     // "min_pow_difficulty": 30,
        // "auth_required": false,
    //     "payment_required": false,
    //     "restricted_writes": false,
    //     // "created_at_lower_limit": 31536000,
    //     // "created_at_upper_limit": 3
    // },
    // "retention": [
    //   {"kinds": [0, 1, [5, 7], [40, 49]], "time": 3600},
    //   {"kinds": [[40000, 49999]], "time": 100},
    //   {"kinds": [[30000, 39999]], "count": 1000},
    //   {"time": 3600, "count": 10000}
    // ],

    "payments_url": "https://anzenkodo.github.io/#support",
    // "fees": {
    //     "admission": [{ "amount": 1000000, "unit": "msats" }],
    //     "subscription": [{ "amount": 5000000, "unit": "msats", "period": 2592000 ],
    //     "publication": [{ "kinds": [4], "amount": 100, "unit": "msats" }],
    // },

    "version": engineInfo.version,
    "software": engineInfo.name,
}
