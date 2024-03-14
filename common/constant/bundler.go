package constant

// https://www.mev.to/builders
type BundlerEndpoint string

const (
	Builder0x69      BundlerEndpoint = "https://builder0x69.io/"
	BeaverBuild      BundlerEndpoint = "https://rpc.beaverbuild.org/"
	FlashbotsBundler BundlerEndpoint = "https://relay.flashbots.net/"
	RsyncBundler     BundlerEndpoint = "https://rsync-builder.xyz/"
	// Blocknative      BundlerEndpoint = "https://api.blocknative.com/v1/auction/"
	GambitLabs      BundlerEndpoint = "https://builder.gmbit.co/rpc/"
	EthBuilder      BundlerEndpoint = "https://eth-builder.com/"
	TitanEndpoint   BundlerEndpoint = "https://rpc.titanbuilder.xyz/"
	BuildAI         BundlerEndpoint = "https://buildai.net/"
	Payload         BundlerEndpoint = "https://rpc.payload.de/"
	Bloxroute       BundlerEndpoint = "https://mev.api.blxrbdn.com/"
	Lightspeed      BundlerEndpoint = "https://rpc.lightspeedbuilder.info/"
	NFactorial      BundlerEndpoint = "https://rpc.nfactorial.xyz/"
	BobaBuilder     BundlerEndpoint = "https://boba-builder.com/searcher/bundle"
	F1b             BundlerEndpoint = "https://rpc.f1b.io/"
	JetBldr         BundlerEndpoint = "https://rpc.jetbldr.xyz/"
	PenguinBuild    BundlerEndpoint = "https://rpc.penguinbuild.org/"
	LokiBuilder     BundlerEndpoint = "https://rpc.lokibuilder.xyz/"
	EdenNetwork     BundlerEndpoint = "https://api.edennetwork.io/v1/bundle/"
	TBuilder        BundlerEndpoint = "https://rpc.tbuilder.xyz/"
	Eigenphi        BundlerEndpoint = "https://builder.eigenphi.io/"
	BlockBeelder    BundlerEndpoint = "https://blockbeelder.com/rpc/"
	ManifoldFinance BundlerEndpoint = "https://api.securerpc.com/v1/"
	PandaBuilder    BundlerEndpoint = "https://rpc.pandabuilder.io/"
	SmithBot        BundlerEndpoint = "https://smithbot.xyz/"
)

func (b BundlerEndpoint) String() string {
	return string(b)
}

func AllBundlerEndpoints() []BundlerEndpoint {
	return []BundlerEndpoint{
		Builder0x69,
		BeaverBuild,
		FlashbotsBundler,
		RsyncBundler,
		// Blocknative,
		GambitLabs,
		EthBuilder,
		TitanEndpoint,
		BuildAI,
		Payload,
		Bloxroute,
		Lightspeed,
		NFactorial,
		BobaBuilder,
		F1b,
		JetBldr,
		PenguinBuild,
		LokiBuilder,
		EdenNetwork,
		TBuilder,
		Eigenphi,
		BlockBeelder,
		ManifoldFinance,
		PandaBuilder,
		SmithBot,
	}
}

type GoerliTestnetBundlerEndpoint string

const (
	FlashbotsGoerliBundler   GoerliTestnetBundlerEndpoint = "https://relay-goerli.flashbots.net/"
	BuildAIGoerliBundler     GoerliTestnetBundlerEndpoint = "https://buildai.net/goerli/"
	EdenNetworkGoerliBundler GoerliTestnetBundlerEndpoint = "https://goerli.edennetwork.io/v1/bundle/"
)

func (t GoerliTestnetBundlerEndpoint) String() string {
	return string(t)
}

func AllGoerliBundlerEndpoints() []GoerliTestnetBundlerEndpoint {
	return []GoerliTestnetBundlerEndpoint{
		FlashbotsGoerliBundler,
		BuildAIGoerliBundler,
		EdenNetworkGoerliBundler,
	}
}

type SepoliaTestnetBundlerEndpoint string

const (
	FlashbotsSepoliaBundler SepoliaTestnetBundlerEndpoint = "https://relay-sepolia.flashbots.net"
)

func AllSepoliaBundlerEndpoints() []SepoliaTestnetBundlerEndpoint {
	return []SepoliaTestnetBundlerEndpoint{
		FlashbotsSepoliaBundler,
	}
}
