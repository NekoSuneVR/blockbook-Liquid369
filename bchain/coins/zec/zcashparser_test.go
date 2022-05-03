//go:build unittest

package zec

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

var (
	testTx1, testTx2 bchain.Tx

	testTxPacked1 = "0a20e64aac0c211ad210c90934f06b1cc932327329e41a9f70c6eb76f79ef798b7b812ab1002000000019c012650c99d0ef761e863dbb966babf2cb7a7a2b5d90b1461c09521c473d23d000000006b483045022100f220f48c5267ef92a1e7a4d3b44fe9d97cce76eeba2785d45a0e2620b70e8d7302205640bc39e197ce19d95a98a3239af0f208ca289c067f80c97d8e411e61da5dee0121021721e83315fb5282f1d9d2a11892322df589bccd9cef45517b5fb3cfd3055c83ffffffff018eec1a3c040000001976a9149bb8229741305d8316ba3ca6a8d20740ce33c24188ac000000000162b4fc6b0000000000000000000000006ffa88c89b74f0f82e24744296845a0d0113b132ff5dfc2af34e6418eb15206af53078c4dd475cf143cd9a427983f5993622464b53e3a37d2519a946492c3977e30f0866550b9097222993a439a39260ac5e7d36aef38c7fdd1df3035a2d5817a9c20526e38f52f822d4db9d2f0156c4119d786d6e3a060ca871df7fae9a5c3a9c921b38ddc6414b13d16aa807389c68016e54bd6a9eb3b23a6bc7bf152e6dba15e9ec36f95dab15ad8f4a92a9d0309bbd930ef24bb7247bf534065c1e2f5b42e2c80eb59f48b4da6ec522319e065f8c4e463f95cc7fcad8d7ee91608e3c0ffcaa44129ba2d2da45d9a413919eca41af29faaf806a3eeb823e5a6c51afb1ec709505d812c0306bd76061a0a62d207355ad44d1ffce2b9e1dfd0818f79bd0f8e4031116b71fee2488484f17818b80532865773166cd389929e8409bb94e3948bd2e0215ef96d4e29d094590fda0de50715c11ff47c03380bb1d31b14e5b4ad8a372ca0b03364ef85f086b8a8eb5c56c3b1aee33e2cfbf1b2be1a3fb41b14b2c432b5d04d54c058fa87a96ae1d65d61b79360d09acc1e25a883fd7ae9a2a734a03362903021401c243173e1050b5cdb459b9ffc07c95e920f026618952d3a800b2e47e03b902084aed7ee8466a65d34abdbbd292781564dcd9b7440029d48c2640ebc196d4b40217f2872c1d0c1c9c2abf1147d6a5a9501895bc92960bfa182ceeb76a658224f1022bc53c4c1cd6888d72a152dc1aec5ba8a1d750fb7e498bee844d3481e4b4cd210227f94f775744185c9f24571b7df0c1c694cb2d3e4e9b955ed0b1caad2b02b5702139c4fbba03f0e422b2f3e4fc822b4f58baf32e7cd217cdbdec8540cb13d6496f271959b72a05e130eeffbe5b9a7fcd2793347cd9c0ea695265669844c363190f690c52a600cf413c3f00bdc5e9d1539e0cc63f4ec2945e0d86e6304a6deb5651e73eac21add5a641dfc95ab56200ed40d81f76755aee4659334c17ed3841ca5a5ab22f923956be1d264be2b485a0de55404510ece5c73d6626798be688f9dc18b69846acfe897a357cc4afe31f57fea32896717f124290e68f36f849fa6ecf76e02087f8c19dbc566135d7fa2daca2d843b9cc5bc3897d35f1de7d174f6407658f4a3706c12cea53d880b4d8c4d45b3f0d210214f815be49a664021a4a44b4a63e06a41d76b46f9aa6bad248e8d1a974ae7bbae5ea8ac269447db91637a19346729083cad5aebd5ff43ea13d04783068e9136da321b1152c666d2995d0ca06b26541deac62f4ef91f0e4af445b18a5c2a17c96eada0b27f85bb26dfb8f16515114c6b9f88037e2b85b3b84b65822eb99c992d99d12dcf9c71e5b46a586016faf5758483a716566db95b42187c101df68ca0554824e1c23cf0302bea03ad0a146af57e91794a268b8c82d78211718c8b5fea286f5de72fc7dfffecddcc02413525c472cb26022641d4bec2b8b7e71a7beb9ee18b82632799498eeee9a351cb9431a8d1906d5164acdf351bd538c3e9d1da8a211fe1cd18c44e72d8cdf16ce3fc9551552c05d52846ea7ef619232102588395cc2bcce509a4e7f150262a76c15475496c923dfce6bfc05871467ee7c213b39ea365c010083e0b1ba8926d3a9e586d8b11c9bab2a47d888bc7cb1a226c0086a1530e295d0047547006f4c8f1c24cdd8e16bb3845749895dec95f03fcda97d3224f6875b1b7b1c819d2fd35dd30968a3c82bc480d10082caf9d9dda8f9ec649c136c7fa07978099d97eaf4abfdc9854c266979d3cfc868f60689b6e3098b6c52a21796fe7c259d9a0dadf1b6efa59297d4c8c902febe7acf826eed30d40d2ac5119be91b51f4839d94599872c9a93c3e2691294914034001d3a278cb4a84d4ae048c0201a97e4cf1341ee663a162f5b586355018b9e5e30624ccdbeacf7d0382afacaf45f08e84d30c50bcd4e55c3138377261deb4e8c2931cd3c51cee94a048ae4839517b6e6537a5c0148d3830a33fea719ef9b4fa437e4d5fecdb646397c19ee56a0973c362a81803895cdc67246352dc566689cb203f9ebda900a5537bbb75aa25ddf3d4ab87b88737a58d760e1d271f08265daae1fe056e71971a8b826e5b215a05b71f99315b167dd2ec78874189657acafac2b5eeb9a901913f55f7ab69e1f9b203504448d414e71098b932a2309db57257eb3fef9de2f2a5a69aa46747d7b827df838345d38b95772bdab8c178c45777b92e8773864964b8e12ae29dbc1b21bf6527589f6bec71ff1cbb9928477409811c2e8150c79c3f21027ee954863b716875d3e9adfc6fdb18cd57a49bb395ca5c42da56f3beb78aad3a7a487de34a870bca61f3cdec422061328c83c910ab32ea7403c354915b7ebee29e1fea5a75158197e4a68e103f017fd7de5a70148ee7ce59356b1a74f83492e14faaa6cd4870bcc004e6eb0114d3429b74ea98fe2851b4553467a7660074e69b040aa31220d0e405d9166dbaf15e3ae2d8ec3b049ed99d17e0743bb6a1a7c3890bbdb7117f7374ad7a59aa1ab47d10445b28f4bc033794a71f88a8bf024189e9d27f9dc5859a4296437585b215656f807aca9dad35747494a43b8a1cf38be2b18a13de32a262ab29f9ba271c4fbce1a470a8243ebf9e7fd37b09262314afbb9a7e180218a0f1c9d505200028b0eb113299010a0012203dd273c42195c061140bd9b5a2a7b72cbfba66b9db63e861f70e9dc95026019c1800226b483045022100f220f48c5267ef92a1e7a4d3b44fe9d97cce76eeba2785d45a0e2620b70e8d7302205640bc39e197ce19d95a98a3239af0f208ca289c067f80c97d8e411e61da5dee0121021721e83315fb5282f1d9d2a11892322df589bccd9cef45517b5fb3cfd3055c8328ffffffff0f3a490a05043c1aec8e10001a1976a9149bb8229741305d8316ba3ca6a8d20740ce33c24188ac222374315934794c31344143486141626a656d6b647057376e594e48576e763179516244414000"
	testTxPacked2 = "0a20bb47a9dd926de63e9d4f8dac58c3f63f4a079569ed3b80e932274a80f60e58b512e20101000000019cafb5c287980e6e5afb47339f6c1c81136d8255f5bd5226b36b01288494c46f000000006b483045022100c92b2f3c54918fa26288530c63a58197ea4974e5b6d92db792dd9717e6d9183c02204e577254213675466a6adad3ae6e9384cf8269fb2dd9943b86fac0c0ad8e3f98012102c99dab469e63b232488b3e7acb9cfcab7e5755f61aad318d9e06b38e5ea22880feffffff0223a7a784010000001976a914826f87806ddd4643730be99b41c98acc379e83db88ac80969800000000001976a914e395634b7684289285926d4c64db395b783720ec88ac6e75040018e4b1c9d50520eeea1128f9ea113299010a0012206fc4948428016bb32652bdf555826d13811c6c9f3347fb5a6e0e9887c2b5af9c1800226b483045022100c92b2f3c54918fa26288530c63a58197ea4974e5b6d92db792dd9717e6d9183c02204e577254213675466a6adad3ae6e9384cf8269fb2dd9943b86fac0c0ad8e3f98012102c99dab469e63b232488b3e7acb9cfcab7e5755f61aad318d9e06b38e5ea2288028feffffff0f3a490a050184a7a72310001a1976a914826f87806ddd4643730be99b41c98acc379e83db88ac22237431566d4854547770457477766f6a786f644e32435351714c596931687a59336341713a470a0398968010011a1976a914e395634b7684289285926d4c64db395b783720ec88ac222374316563784d587070685554525158474c586e56684a367563714433445a69706464674000"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "02000000019c012650c99d0ef761e863dbb966babf2cb7a7a2b5d90b1461c09521c473d23d000000006b483045022100f220f48c5267ef92a1e7a4d3b44fe9d97cce76eeba2785d45a0e2620b70e8d7302205640bc39e197ce19d95a98a3239af0f208ca289c067f80c97d8e411e61da5dee0121021721e83315fb5282f1d9d2a11892322df589bccd9cef45517b5fb3cfd3055c83ffffffff018eec1a3c040000001976a9149bb8229741305d8316ba3ca6a8d20740ce33c24188ac000000000162b4fc6b0000000000000000000000006ffa88c89b74f0f82e24744296845a0d0113b132ff5dfc2af34e6418eb15206af53078c4dd475cf143cd9a427983f5993622464b53e3a37d2519a946492c3977e30f0866550b9097222993a439a39260ac5e7d36aef38c7fdd1df3035a2d5817a9c20526e38f52f822d4db9d2f0156c4119d786d6e3a060ca871df7fae9a5c3a9c921b38ddc6414b13d16aa807389c68016e54bd6a9eb3b23a6bc7bf152e6dba15e9ec36f95dab15ad8f4a92a9d0309bbd930ef24bb7247bf534065c1e2f5b42e2c80eb59f48b4da6ec522319e065f8c4e463f95cc7fcad8d7ee91608e3c0ffcaa44129ba2d2da45d9a413919eca41af29faaf806a3eeb823e5a6c51afb1ec709505d812c0306bd76061a0a62d207355ad44d1ffce2b9e1dfd0818f79bd0f8e4031116b71fee2488484f17818b80532865773166cd389929e8409bb94e3948bd2e0215ef96d4e29d094590fda0de50715c11ff47c03380bb1d31b14e5b4ad8a372ca0b03364ef85f086b8a8eb5c56c3b1aee33e2cfbf1b2be1a3fb41b14b2c432b5d04d54c058fa87a96ae1d65d61b79360d09acc1e25a883fd7ae9a2a734a03362903021401c243173e1050b5cdb459b9ffc07c95e920f026618952d3a800b2e47e03b902084aed7ee8466a65d34abdbbd292781564dcd9b7440029d48c2640ebc196d4b40217f2872c1d0c1c9c2abf1147d6a5a9501895bc92960bfa182ceeb76a658224f1022bc53c4c1cd6888d72a152dc1aec5ba8a1d750fb7e498bee844d3481e4b4cd210227f94f775744185c9f24571b7df0c1c694cb2d3e4e9b955ed0b1caad2b02b5702139c4fbba03f0e422b2f3e4fc822b4f58baf32e7cd217cdbdec8540cb13d6496f271959b72a05e130eeffbe5b9a7fcd2793347cd9c0ea695265669844c363190f690c52a600cf413c3f00bdc5e9d1539e0cc63f4ec2945e0d86e6304a6deb5651e73eac21add5a641dfc95ab56200ed40d81f76755aee4659334c17ed3841ca5a5ab22f923956be1d264be2b485a0de55404510ece5c73d6626798be688f9dc18b69846acfe897a357cc4afe31f57fea32896717f124290e68f36f849fa6ecf76e02087f8c19dbc566135d7fa2daca2d843b9cc5bc3897d35f1de7d174f6407658f4a3706c12cea53d880b4d8c4d45b3f0d210214f815be49a664021a4a44b4a63e06a41d76b46f9aa6bad248e8d1a974ae7bbae5ea8ac269447db91637a19346729083cad5aebd5ff43ea13d04783068e9136da321b1152c666d2995d0ca06b26541deac62f4ef91f0e4af445b18a5c2a17c96eada0b27f85bb26dfb8f16515114c6b9f88037e2b85b3b84b65822eb99c992d99d12dcf9c71e5b46a586016faf5758483a716566db95b42187c101df68ca0554824e1c23cf0302bea03ad0a146af57e91794a268b8c82d78211718c8b5fea286f5de72fc7dfffecddcc02413525c472cb26022641d4bec2b8b7e71a7beb9ee18b82632799498eeee9a351cb9431a8d1906d5164acdf351bd538c3e9d1da8a211fe1cd18c44e72d8cdf16ce3fc9551552c05d52846ea7ef619232102588395cc2bcce509a4e7f150262a76c15475496c923dfce6bfc05871467ee7c213b39ea365c010083e0b1ba8926d3a9e586d8b11c9bab2a47d888bc7cb1a226c0086a1530e295d0047547006f4c8f1c24cdd8e16bb3845749895dec95f03fcda97d3224f6875b1b7b1c819d2fd35dd30968a3c82bc480d10082caf9d9dda8f9ec649c136c7fa07978099d97eaf4abfdc9854c266979d3cfc868f60689b6e3098b6c52a21796fe7c259d9a0dadf1b6efa59297d4c8c902febe7acf826eed30d40d2ac5119be91b51f4839d94599872c9a93c3e2691294914034001d3a278cb4a84d4ae048c0201a97e4cf1341ee663a162f5b586355018b9e5e30624ccdbeacf7d0382afacaf45f08e84d30c50bcd4e55c3138377261deb4e8c2931cd3c51cee94a048ae4839517b6e6537a5c0148d3830a33fea719ef9b4fa437e4d5fecdb646397c19ee56a0973c362a81803895cdc67246352dc566689cb203f9ebda900a5537bbb75aa25ddf3d4ab87b88737a58d760e1d271f08265daae1fe056e71971a8b826e5b215a05b71f99315b167dd2ec78874189657acafac2b5eeb9a901913f55f7ab69e1f9b203504448d414e71098b932a2309db57257eb3fef9de2f2a5a69aa46747d7b827df838345d38b95772bdab8c178c45777b92e8773864964b8e12ae29dbc1b21bf6527589f6bec71ff1cbb9928477409811c2e8150c79c3f21027ee954863b716875d3e9adfc6fdb18cd57a49bb395ca5c42da56f3beb78aad3a7a487de34a870bca61f3cdec422061328c83c910ab32ea7403c354915b7ebee29e1fea5a75158197e4a68e103f017fd7de5a70148ee7ce59356b1a74f83492e14faaa6cd4870bcc004e6eb0114d3429b74ea98fe2851b4553467a7660074e69b040aa31220d0e405d9166dbaf15e3ae2d8ec3b049ed99d17e0743bb6a1a7c3890bbdb7117f7374ad7a59aa1ab47d10445b28f4bc033794a71f88a8bf024189e9d27f9dc5859a4296437585b215656f807aca9dad35747494a43b8a1cf38be2b18a13de32a262ab29f9ba271c4fbce1a470a8243ebf9e7fd37b09262314afbb9a7e1802",
		Blocktime: 1521645728,
		Time:      1521645728,
		Txid:      "e64aac0c211ad210c90934f06b1cc932327329e41a9f70c6eb76f79ef798b7b8",
		LockTime:  0,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100f220f48c5267ef92a1e7a4d3b44fe9d97cce76eeba2785d45a0e2620b70e8d7302205640bc39e197ce19d95a98a3239af0f208ca289c067f80c97d8e411e61da5dee0121021721e83315fb5282f1d9d2a11892322df589bccd9cef45517b5fb3cfd3055c83",
				},
				Txid:     "3dd273c42195c061140bd9b5a2a7b72cbfba66b9db63e861f70e9dc95026019c",
				Vout:     0,
				Sequence: 4294967295,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(18188266638),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a9149bb8229741305d8316ba3ca6a8d20740ce33c24188ac",
					Addresses: []string{
						"t1Y4yL14ACHaAbjemkdpW7nYNHWnv1yQbDA",
					},
				},
			},
		},
	}

	testTx2 = bchain.Tx{
		Hex:       "01000000019cafb5c287980e6e5afb47339f6c1c81136d8255f5bd5226b36b01288494c46f000000006b483045022100c92b2f3c54918fa26288530c63a58197ea4974e5b6d92db792dd9717e6d9183c02204e577254213675466a6adad3ae6e9384cf8269fb2dd9943b86fac0c0ad8e3f98012102c99dab469e63b232488b3e7acb9cfcab7e5755f61aad318d9e06b38e5ea22880feffffff0223a7a784010000001976a914826f87806ddd4643730be99b41c98acc379e83db88ac80969800000000001976a914e395634b7684289285926d4c64db395b783720ec88ac6e750400",
		Blocktime: 1521637604,
		Time:      1521637604,
		Txid:      "bb47a9dd926de63e9d4f8dac58c3f63f4a079569ed3b80e932274a80f60e58b5",
		LockTime:  292206,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "483045022100c92b2f3c54918fa26288530c63a58197ea4974e5b6d92db792dd9717e6d9183c02204e577254213675466a6adad3ae6e9384cf8269fb2dd9943b86fac0c0ad8e3f98012102c99dab469e63b232488b3e7acb9cfcab7e5755f61aad318d9e06b38e5ea22880",
				},
				Txid:     "6fc4948428016bb32652bdf555826d13811c6c9f3347fb5a6e0e9887c2b5af9c",
				Vout:     0,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(6520547107),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914826f87806ddd4643730be99b41c98acc379e83db88ac",
					Addresses: []string{
						"t1VmHTTwpEtwvojxodN2CSQqLYi1hzY3cAq",
					},
				},
			},
			{
				ValueSat: *big.NewInt(10000000),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914e395634b7684289285926d4c64db395b783720ec88ac",
					Addresses: []string{
						"t1ecxMXpphUTRQXGLXnVhJ6ucqD3DZipddg",
					},
				},
			},
		},
	}
}

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func TestGetAddrDesc(t *testing.T) {
	type args struct {
		tx     bchain.Tx
		parser *ZCashParser
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "zec-1",
			args: args{
				tx:     testTx1,
				parser: NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
		},
		{
			name: "zec-2",
			args: args{
				tx:     testTx2,
				parser: NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for n, vout := range tt.args.tx.Vout {
				got1, err := tt.args.parser.GetAddrDescFromVout(&vout)
				if err != nil {
					t.Errorf("getAddrDescFromVout() error = %v, vout = %d", err, n)
					return
				}
				got2, err := tt.args.parser.GetAddrDescFromAddress(vout.ScriptPubKey.Addresses[0])
				if err != nil {
					t.Errorf("getAddrDescFromAddress() error = %v, vout = %d", err, n)
					return
				}
				if !bytes.Equal(got1, got2) {
					t.Errorf("Address descriptors mismatch: got1 = %v, got2 = %v", got1, got2)
				}
			}
		})
	}
}

func TestPackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *ZCashParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "zec-1",
			args: args{
				tx:        testTx1,
				height:    292272,
				blockTime: 1521645728,
				parser:    NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
		{
			name: "zec-2",
			args: args{
				tx:        testTx2,
				height:    292217,
				blockTime: 1521637604,
				parser:    NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    testTxPacked2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func TestUnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *ZCashParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "zec-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   292272,
			wantErr: false,
		},
		{
			name: "zec-2",
			args: args{
				packedTx: testTxPacked2,
				parser:   NewZCashParser(GetChainParams("main"), &btc.Configuration{}),
			},
			want:    &testTx2,
			want1:   292217,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
