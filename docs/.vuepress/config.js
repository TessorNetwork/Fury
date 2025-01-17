module.exports = {
    title: "Tessor.network Documentation",
    description: "Documentation for the Tessor.network blockchain.",
    ga: "UA-51029217-2",
    head: [
        ['link', {rel: 'icon', href: '/icon.png'}],
    ],
    /*markdown: {
        lineNumbers: true,
	    extendMarkdown: md => {
		    md.use(require("markdown-it-footnote"));
  	    }
    },*/
    plugins: [
        '@renovamen/vuepress-plugin-katex',
        [
          "sitemap",
          {
            hostname: "https://docs.tessor.network"
          }
        ]
    ],
    themeConfig: {
        repo: "tessornetwork/fury",
        logo: '/icon.png',
        editLinks: true,
        docsDir: "docs",
        docsBranch: "master",
        editLinkText: 'Edit this page on Github',
        lastUpdated: true,
        nav: [
            {
                text: 'Versions',
                ariaLabel: 'Versions',
                items: [
                    { text: 'ver 3.0.0', link: '/' },
                    { text: 'ver 2.2.0', link: '/docs2.2.0/' },
                    { text: 'ver 2.1.2', link: '/docs2.1.2/' }
                ]
            },
            {text: "Tessor.network", link: "https://tessor.network"},
            {text: "Discord", link: "https://discord.gg/N7DxaDj5sW"},
            {text: "Telegram", link: "https://t.me/TessorNetwork"},
            {text: "Twitter", link: "https://twitter.com/furynet?s=21&t=8FTpg5f7kurZ1d7LOb9YXw"},
        ],

        sidebarDepth: 3,
        sidebar: [
            {
                title: "Overview",
                collapsable: false,
                path: "/"
            },
            {
                title: "Running Nodes",
                collapsable: false,
                children: [
                    ["nodes/", "Tessor.network overview"],
                    ["nodes/hardware-requirements", "Hardware requirements"],
                    ["nodes/full-node-installation", "Installing a full node"],
                    ["nodes/validator-node-installation", "Becoming a validator"],
                    ["nodes/validator-node-handling", "Handling a validator"],
                    ["nodes/validator-node-update", "Updating a validator"]
                ]
            },
            {
                title: "API Developers",
                collapsable: false,
                children: [
                    ["app_developers/furyapi-introduction", "Introduction to FuryAPI"],
                    ["app_developers/furyapi-authentication", "Authentication process"],
                    ["app_developers/furyapi-wallet", "Wallet"],
                    ["app_developers/furyapi-sharedoc", "ShareDoc"],

                ]
            },
            {
                title: "Custom Modules",
                path: "/modules/",
                collapsable: false,
                children: [
                    //["x/bank/","Bank"],
                    {
                        title: "Government",
                        path: "/modules/government/",
                        collapsable: true,
                        children: [
                            ["modules/government/", "Concepts"],
                            ["modules/government/01_state.md", "State"],
                        ]
                    },                       
                    {
                        title: "Did",
                        path: "/modules/did/",
                        collapsable: true,
                        children: [
                            ["modules/did/", "Concepts"],
                            ["modules/did/01_state.md", "State"],
                            ["modules/did/05_client.md", "Client"],

                        ]
                    },                    
                    {
                        title: "Documents",
                        path: "/modules/documents/",
                        collapsable: true,
                        children: [
                            ["modules/documents/", "Concepts"],
                            ["modules/documents/01_state.md", "State"],
                            //["modules/documents/02_keepers.md", "Keepers"],
                            ["modules/documents/03_messages.md", "Messages"],
                            ["modules/documents/04_events.md", "Events"],
                            ["modules/documents/05_client.md", "Client"],

                        ]
                    },
                    {
                        title: "FuryMint",
                        path: "/modules/furymint/",
                        collapsable: true,
                        children: [
                            ["modules/furymint/", "Concepts"],
                            ["modules/furymint/01_state.md", "State"],
                            ["modules/furymint/02_messages.md", "Messages"],
                            ["modules/furymint/03_events.md", "Events"],
                            ["modules/furymint/04_params.md", "Params"],
                            ["modules/furymint/05_client.md", "Client"],

                        ]
                    },
                    {
                        title: "FuryKYC",
                        path: "/modules/furykyc/",
                        collapsable: true,
                        children: [
                            ["modules/furykyc/", "Concepts"],
                            ["modules/furykyc/01_state.md", "State"],
                            ["modules/furykyc/02_messages.md", "Messages"],
                            ["modules/furykyc/03_events.md", "Events"],
                            ["modules/furykyc/04_client.md", "Client"],

                        ]
                    },
                    {
                        title: "Vbr",
                        path: "/modules/vbr/",
                        collapsable: true,
                        children: [
                            ["modules/vbr/", "Concepts"],
                            ["modules/vbr/01_state.md", "State"],
                            ["modules/vbr/02_messages.md", "Messages"],
                            ["modules/vbr/03_events.md", "Events"],
                            ["modules/vbr/04_params.md", "Params"],
                            ["modules/vbr/05_client.md", "Client"],

                        ]
                    },

                ]
            },
            {
                title: "ver 2.2.0",
                collapsable: true,
                children: [
                    ["docs2.2.0/", "ver 2.2.0"],
                    {
                        title: "Nodes",
                        collapsable: true,
                        children: [
                            ["docs2.2.0/nodes/", "Introduction"],
                            ["docs2.2.0/nodes/hardware-requirements", "Hardware requirements"],
                            ["docs2.2.0/nodes/full-node-installation", "Installing a full node"],
                            ["docs2.2.0/nodes/validator-node-installation", "Becoming a validator"],
                            ["docs2.2.0/nodes/validator-node-handling", "Handling a validator"],
                            ["docs2.2.0/nodes/validator-node-update", "Updating a validator"],
                        ]
                    },
                    {
                        title: "App Developers",
                        collapsable: true,
                        children: [
                            ["docs2.2.0/app_developers/", "Introduction"]
                        ]
                    },
                    {
                        title: "SDK Developers",
                        collapsable: true,
                        children: [
                            ["docs2.2.0/developers/", "Introduction"],
                            "docs2.2.0/developers/create-sign-broadcast-tx",
                            "docs2.2.0/developers/message-types",
                            "docs2.2.0/developers/listing-transactions"
                        ]
                    },


                    {
                        title: "Modules",
                        collapsable: true,
                        children: [
                            "docs2.2.0/x/bank/",
                            "docs2.2.0/x/government/",
                            "docs2.2.0/x/id/",
                            "docs2.2.0/x/documents/",
                            "docs2.2.0/x/furymint/",
                            "docs2.2.0/x/furykyc/",
                            "docs2.2.0/x/vbr/",
                        ]
                    }
                ]
            },
            {
                title: "ver 2.1.2",
                collapsable: true,
                children: [
                    ["docs2.1.2/", "ver 2.1.2"],
                    {
                        title: "Nodes",
                        collapsable: true,
                        children: [
                            ["docs2.1.2/nodes/", "Introduction"],
                            ["docs2.1.2/nodes/hardware-requirements", "Hardware requirements"],
                            ["docs2.1.2/nodes/full-node-installation", "Installing a full node"],
                            ["docs2.1.2/nodes/validator-node-installation", "Becoming a validator"],
                            ["docs2.1.2/nodes/validator-node-handling", "Handling a validator"],
                            ["docs2.1.2/nodes/validator-node-update", "Updating a validator"],
                        ]
                    },
                    {
                        title: "App Developers",
                        collapsable: true,
                        children: [
                            ["docs2.1.2/app_developers/", "Introduction"]
                        ]
                    },
                    {
                        title: "SDK Developers",
                        collapsable: true,
                        children: [
                            ["docs2.1.2/developers/", "Introduction"],
                            "docs2.1.2/developers/create-sign-broadcast-tx",
                            "docs2.1.2/developers/message-types",
                            "docs2.1.2/developers/listing-transactions"
                        ]
                    },


                    {
                        title: "Modules",
                        collapsable: true,
                        children: [
                            "docs2.1.2/x/bank/",
                            "docs2.1.2/x/government/",
                            "docs2.1.2/x/id/",
                            "docs2.1.2/x/docs/",
                            "docs2.1.2/x/pricefeed/",
                            "docs2.1.2/x/furymint/",
                            "docs2.1.2/x/memberships/",
                            "docs2.1.2/x/vbr/",
                            "docs2.1.2/x/creditrisk/"
                        ]
                    }
                ]
            }
        ],
    }
};
