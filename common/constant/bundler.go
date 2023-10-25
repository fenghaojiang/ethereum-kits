package constant

// https://www.mev.to/builders
type BundlerEndpoint string

const (
	Builder0x69      BundlerEndpoint = "https://builder0x69.io/"
	BeaverBuild      BundlerEndpoint = "https://rpc.beaverbuild.com/"
	FlashbotsBundler BundlerEndpoint = "https://relay.flashbots.net/"
	RsyncBundler     BundlerEndpoint = "https://rsync-builder.xyz/"
	Blocknative      BundlerEndpoint = "https://api.blocknative.com/v1/auction/"
	GambitLabs       BundlerEndpoint = "https://builder.gmbit.co/rpc/"
	EthBuilder       BundlerEndpoint = "https://eth-builder.com/"
	TitanEndpoint    BundlerEndpoint = "https://rpc.titanbuilder.xyz/"
	BuildAI          BundlerEndpoint = "https://buildai.net"
	Payload          BundlerEndpoint = "https://rpc.payload.de/"
	Bloxroute        BundlerEndpoint = "https://mev.api.blxrbdn.com/"
	Lightspeed       BundlerEndpoint = "https://rpc.lightspeedbuilder.info/"
)

type TestnetBundlerEndpoint string

const (
	FlashbotsGoerliBundler TestnetBundlerEndpoint = "https://relay-goerli.flashbots.net/"
)

func (t TestnetBundlerEndpoint) String() string {
	return string(t)
}