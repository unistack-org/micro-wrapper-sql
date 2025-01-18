// Code generated. DO NOT EDIT.

package wrapper

import "database/sql/driver"

func wrapConn(dc driver.Conn, opts Options) driver.Conn {
	c := &wrapperConn{conn: dc, opts: opts}
	if _, ok := dc.(wrapConn0008_fd07920044cbf1b6fcf3603ebb166c6c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_13fbd5c85aa99a91119179870858fc36); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_d69fa1ea920d7f3f8644d49fc0b6aded); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_86034fd5bf9a19faed074c390256416d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_8a643673abce113f9fea5ce1702250a6); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_265661e823b3aed3ebfb5ca65c6bba29); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_8b66ea616b172e9ee49fd7acdf30b938); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_dd667cce68b8575b61aaaa28533db684); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0007_5b4064ed641abac0e70005efb3a4ced2); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_5046950504d8a3b58a4241a8d69a199f); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_a78ba5b57311d51ec885c88fcf62a0d3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_f476294ed42363d7c2356114aab61ba3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_7a5fa0f486ff80bb9c7fb0e446b45447); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_6ebbdc3a0d909a8ad477a64e2c788e4f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_db09fdc7c55f1435e80abe7fc49f0851); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_976d373277dceea1dca120722cddeb69); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_8c4a4d51c5e2a5272c519b0c73f39980); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_4139e5345038cc13f9683ec9a592592c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_e25c3a362b3347dd39e21d3465ef0049); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_34d5ca3511ee070149519075ad8a1d35); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_bbd5fd8057adff78016752e1be0a9d10); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_62f57ede9f4fdca4008bb8b102bff7ba); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_f696e01b39fbfe1dd45edc9adeb4154a); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_6bb429c97c0ecdbf2ad34c4ca1332c21); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_aa58b7d423443aa26503c4665dc3743d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_4186142f3f5124cf0aa4c9fcf856b2c2); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_fdb81facd138c4a0705bd634d6c5ec20); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_ec5f3f3ec9a278b63dda59c8c26ba82e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_5802b2f9aac2139509d07fe2b7e4c96b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_b023136fe1c19ee81c3975c762b0b8e8); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_21c6aaf05ebd245d2a81458a7c024555); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_d079c9555b86aa5a077b492b67851ebc); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_59a3746398209af8d5abd1b9cc6e3ec8); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_8f10b08e9d138b1d7c3584d1ecc1451e); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_6f5a532ec7684dbd100cad131c0e6938); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_a0bfa772ad0992f8541f658be51b7492); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0006_8dc061cdbbf61be8784e703979185c75); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_5698322bcb20d1da317fbe9951ea1ea8); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_c670b45682d64d3a2f572ae08687b1c3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_d98ff0a96153b4e2dd5d28a88b553717); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_95dabc404203cf349885a9d084b8cd6d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_036cfac2479cb17c204aeee52931f723); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_aa238e8cc38c595d62af9027950c4b6d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_20631a232115a4d613bbdd90bd786e91); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_bb0588c994e9ade1989ee58ccd0341e0); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_2b9c34cec984ef83aa9d7eb139b39e1e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_58c6e55019354e20f94dc97b5f01710e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_020e8fd0be2f1cfda622cd7fae226094); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_c19e3acd7931b2ef3bd181d17c473d6c); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_13d3220ae21dcb3631a0e0d628adb44b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_455bce9dcc99aa7c0fe96735f9c1c93c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_172f2782b0ad97eecc7f56f17323847c); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_e5655e77364e3f931a3ea6875c01a0e7); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_26c5be906ea66f39030cc40db09bbbef); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_a6894db2a6de9012f66db439c1729127); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_86594d133db62770dc866e4b90a41484); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_be4447c4210552d0e2d709e04422e46c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_9820f7ab194cfc8681fa530ec810fbee); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_20d8d229b2484368b368c970e6deaaec); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_fd14a18f00f97dfb88beee48134b1ed7); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_612cbcd42b1ca988c32b19b0ee60fe8e); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_858ad5997c7a1ecd0c4e016da0b78cb3); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_0ce7c7f80ec4cb8b2af6490440fb797b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_6134bd04f1e6222f7503813d44eae9a4); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_522573917a875da9c8a8afd3f5bc8e45); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_c391b2ea017883e989dd4c91421a2518); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_afdced595d6aac26d33b2396172517d3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_c4039cd70e55947c4038bbc29a3711d7); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_6f9178fa3ebe1aed2696357cd2afb3a9); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_650d3e7f112f289ca18b3e91964181c8); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_2e8be41ad679566886fbb5aa5471154b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_1f8c595ee461d1900fce4cab06c9d1a3); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_a03e3132f8a1d5217b2ffa017ef0e471); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_5e65f4368961d057065044067e0fdfa0); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_61a248bb5bae07d449dffffcab033bd3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_48bc1b8fb59601a4f7543709bef7460b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_8e89c1bd95e4e9731efc7ac1abba3639); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_370f70e1f7779d2b18c16b77796e6ad4); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_dccec1fa0b1ac6076c36fab915923e7c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_db318a810481bb79ec5c2f495823168c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_7a8098e8711012274339764352391092); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_0e4f2ce6ee8c6aaba65ea53cad0a7eda); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_dab2c71afcbd6af7c7e4e507d436ebb3); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_8ea2641f961a5c26a90da3a80a453961); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_65e1c77325da4d0c054e69bd4551c599); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_c829a7b8ed999b05fe10801538e8fc57); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_3dde8f11372ab3df00c06e9cc9f532a0); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_436b3b6f6e649ed483a53237c3775c95); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_497c507b429fdc6223283c0a0a534ea5); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_81570b25b723de4082fb023c40e4ccd9); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_063896be35045b544dfdc8bd6e379acb); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_646f19e9526b1b663e651ef0e2bef0d1); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0005_3c446fe7d7a33ac81540f3f0173a2617); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_a0bae6dc5e6fadde4228e11bf5e49dd1); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_28c597b75ba228e5d6902ddc68ad7811); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_d3598002e745fefcab218230adaf46d1); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_4ff0f4092dd79c548ab0e495e6cddc3b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_4ef020c4936aea427d194c2b96378df5); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_f018b29703e3acb35c78e692c0b4e037); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_0fd04b4e30706876c57056f2a95c5b01); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_c818cf1c81182b6fe0e9f65adf919fcc); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_c8f1b5b5a8743b962a62682d4a6e9262); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_15e77a0e9aeb76ab2f04aa748ac4719e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_86042fa2173f60c1aaa263c2c4f6b7ee); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_645074a289d63811b9e9a8978d5ba260); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_2a24acd339cf38eb1b877ff6e36477c2); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_72837a807cd7df18bbd7e22ee697c816); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_e096a7508a321da18cfa19aba7f60b47); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_62deeb4f7a0c751fdee63cefb834691f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_573d41654139e877dea050658666ffa7); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_84343f2e4ec9d07e1a9dbe14088de505); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_d3d3167450d7d79e6e587b43b5014c17); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_555cdb27b272a3f2c0eb9f7b9e0eb5ff); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_8fb2e6de52e01bddf84ba9aa01f20489); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_052895537635d2f5a24651bcc48ec975); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_f39e7f23bc4a7c1bd04d47bcdf9f2b75); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_8c1f1ab2627a320478eb2bd9c49efd73); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_d1052a967ff8db9233909609d294e184); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_51825a78fc953cc98d20358550621612); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_95c3412bca03d21c7267d3279e2f3312); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_684b88bf9d678a38c20afbfe31f4589a); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_089675b7ca4c9fa833b6515226a0b60c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_4f64f93a4e925dbd2a7598b37bb8f332); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_994bda6d7619bde46104d2eab7ff768c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_bf56dd41a04eebaa9037e34d998af495); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_a602fc8de20f9667dca43dab54118301); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_8fef826c40cffbdc6a69c70b5676c824); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_4501f160909b47a29559b718409520ea); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_87e9112616a76f4e1f0250177acda175); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_5d17fc1b7064520ae6022fd2f62c2bfb); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_65ab36fec2bb30a1596eb25d39a63d3c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_ec4d0075fd0ddc8b390599c00aa63ac1); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_ff5d54172f489033d905c69de9288f35); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_4eafb813d5396d818b628d6737e80d42); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_1666fda37418e8810185898fd16573d2); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_b214fbc2704c5d1b1313515138517b76); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_86f21eab726445da5b1ba4da19a01e24); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_b454cff4daf9ab30e3870da405d5fc28); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_002ba8ae7ab8583e9fe9819c79f7f6e6); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_3b5220a9e11f03a530f5405a48a3583d); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_8480b7a2f4b03ada42be3b2e1e8ee03b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_421d78dfe11851b6277447d6fc3d762b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_9efc0e3a9cf918d9d10523202afb880a); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_f33d7e93149b996e3b7d599a621eb41b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_06ee4205d01722f7549f02e0b7087b0b); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_6afb89b2559180fa651920ad74abbb9c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_5423a435d10b320ffac1aac27fe43c6f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_457719d02ae8d6b7caabc81e5ac31d57); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_c44f3469765232bf2f3748ac3491ee68); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_393cb9100ac7954d3e25afc8842b5dcb); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_09adbe5008d26d1ca9b6390889d3e56d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_6e5b2720e60eff99a0177f056c13773b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_c68aa300799a706219ed2f7ad15362c5); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_1b5b9005d979312b723d3a27c0c65db6); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_7c2de3c593da1d0c0ca997bfb7e1d325); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_c6e51aefd95b2286637499cab5834392); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_25d1043092b1d4fde69e802678864cbe); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_ff336118365ff6fcd2a2e761486717c7); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_1d4652a872051fa3240ca534bef0a43e); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_b0e5c7fb09f6acbdfcf6e0d4258d4a9b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_8098fe71a9a9219d58e590ca89192425); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_6fcd36e5fdd1db57ef48d5bdd4beebd9); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0004_47c38651973c418d8726100d2bdf9d3a); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_23392d2c642abd59125f77ecca6f7b8f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_67e361827a469e5a7fd41b0778aa1958); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_c18fe73278d7622bc6a1b7ff02b795e7); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_9d0152a9df319a9e7e4a0bdc7f7fa2d4); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_75415c9d129d1dd60031d39ea83f5316); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_d40c325b5125e64636b9b8ea3f1b8975); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_0980f76b2d444d923da26fc4a1b82818); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.ExecerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_1a34bf796737439096bfa6f10d5696c2); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.NamedValueChecker		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_3371fc70ba5d416a901c5118687e778e); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ddd509d4dc401e51bac497647969675e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_66e6b00152e5329959d8d66776407655); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ce39b6405b9206d7ba9fe332b19cf89f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ff274753683e1d0adf6ecc55b9be8631); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.NamedValueChecker		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_613592023fc524d8f895420661ca88e0); ok {
	return struct {
		driver.Conn
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ecd9017910d8eaafad4c84916460cde2); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_4afcf6b296cf46480b7f20345361eec9); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_2c310269b09748f7be39d17b8528f2df); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_214d339ea349168fbe3e0955424c5eef); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_4b5d9a137b792a784a0b494e06773edb); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_67a6f30839ce48a2e9f3488bf5e2053d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_1bf343887515a244e3d82298438693bd); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_3046c3f6840014e9d6851479421aa2db); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_429af3d54eff18d5961d7d87b53413b7); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_348becb0e4f2e9da375f40a1e59807d1); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_1c8897aa6698e41ba80853991cc24839); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_c02e55bb6f1875e5ca0c1667579be31a); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.SessionResetter		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_d033831362fd7cd8abd09076a328bf5f); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_5a07d33df77df1576ce8067b8997d413); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_58164b71668ef615a11f9d6dd3ff1e07); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_f3e446fd938f183e46a975476439f43d); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ddca4fe2167f6deb65f9ee0d4d46b23b); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_55276ca2338c0cd130c39fb53ac4caaa); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_344b3b12f6cdf795bbc480ca3c22d0eb); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_68f463fe3d9a985f2c57f5f872c7c333); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_b2d88477280fe09a57ad676180ae7fca); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_aebab2fa61e210e3c96cb924b4356379); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_55b0c9c15fb89c88e545e8b9afe973ee); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_57e946d049a38df9bcfe6ef5d8598cfc); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_4e478d169a0c925294ee737c49c1ac6a); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_065f30b751342fccddaa383211967ee5); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_80b435a335887493ba9656f5a0e76a14); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_975ad90f36f2a8d5d14cbc769a6ce607); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ecd209322f30837e1dde5cc3c5903232); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_f589cedd98232e73dba33f2b1ee57f05); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_895867613f93da07976594c2ea883249); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_6168ba655398488893634accebfceca9); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_7908fe1d05a5c3552e8ea5c1ee7d48e7); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_0040353f748f8b98cb8da172c6b336c1); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_42faf5a4eb464290a69fa9800a551bd4); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_ee7a3b8aa7f629ea0410d3cb5df23364); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_4a1ee3cec4754da65ab631f8d55709d8); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_86f0c8258522256415a6791126d2786c); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_cdbed2933e5c0be5f6806d38207387b1); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_66ae780e5764601ed26339f0ef146ab9); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_e8d66097d62d92ca5e8a01d0627ef47b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0003_086984675e3daca69cb4d93a08b03c05); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator		
}{c, c, c, c}
}

	if _, ok := dc.(wrapConn0002_9998c1796991ed1018aa6832ef10c20d); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.QueryerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_608b01490fc5eeb54196421e44a1d294); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_e5960e40a7bbe1b5a19c38f71e0e6861); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.QueryerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_21490f92aea80b01cad7a6a9b496de9d); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_b7ff310fb7e22a7ccbb847bf26429dd9); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ConnPrepareContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_734bbe2891841d705d0385cb4817474b); ok {
	return struct {
		driver.Conn
			driver.QueryerContext
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_fb43968160809b260946ae6f55b7fe56); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.QueryerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_57475ba369865550e7711feefe48c5d4); ok {
	return struct {
		driver.Conn
			driver.QueryerContext
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_dea76879684fe95f8ca4bec7f7cec2d7); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_9743a8d2379220bf3bef64c2a633bd0d); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_01ce0f63e6a6b920df90cbea6b6ede03); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_f21346c5fd58f22896f02b5b5ed3475e); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_0921fb5e4c916e10370d5e0398435d4e); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.Pinger		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_1bf71fb4be83d1fefe01cdd4d5a7a547); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_35700b3f7c770fcd0bf83836eee99fec); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.NamedValueChecker		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_c4cbf403989171fe81ee42f1529c60fe); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_b0936442b3763765deb78c0f8945bb0a); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Pinger		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_20cb62afffd12ef5a3dd67710522c25c); ok {
	return struct {
		driver.Conn
			driver.SessionResetter
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_039f074dd8b7cd2fd02bcb0f3a89e29a); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Pinger		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_defc836eab973764559fef7a6eacab29); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Pinger		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_f2e0c829a13f0fc5e34c7487f6a5743b); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.ExecerContext		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_739d949041a357f14dd543ae1c82d4b4); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_004b3f5abaa12288a30861da40988df6); ok {
	return struct {
		driver.Conn
			driver.Pinger
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_7eb98af2c595d337fef40bcf00e367f5); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_70a606d71e13e67187acc5d5056da2e1); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_1c65fb5735328219b65d13d4e717b79f); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_26864a2d0c8c5a605e497ec245ca9aa1); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.SessionResetter		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0002_b1c510487a94a905aff86496f31078e6); ok {
	return struct {
		driver.Conn
			driver.ExecerContext
			driver.Validator		
}{c, c, c}
}

	if _, ok := dc.(wrapConn0001_dd0fe84666fad76cbf0dee440e953d06); ok {
	return struct {
		driver.Conn
			driver.NamedValueChecker		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_d815927e8e5ae23caed6edbf82ddb9a8); ok {
	return struct {
		driver.Conn
			driver.ConnBeginTx		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_f95e29e11298b900f8a05d72dcfe041b); ok {
	return struct {
		driver.Conn
			driver.Pinger		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_70145ca1951f1697e67aa095f55b8305); ok {
	return struct {
		driver.Conn
			driver.SessionResetter		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_f98304698b6793d84b0e7be539d135ce); ok {
	return struct {
		driver.Conn
			driver.ExecerContext		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_7a610c95e7c563485e18c228cc42e7fb); ok {
	return struct {
		driver.Conn
			driver.QueryerContext		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_c3128ba459e963c7f07c0de8d58b775f); ok {
	return struct {
		driver.Conn
			driver.ConnPrepareContext		
}{c, c}
}

	if _, ok := dc.(wrapConn0001_abc062388a1f7a907b94470050b8bc86); ok {
	return struct {
		driver.Conn
			driver.Validator		
}{c, c}
}

return c
}
// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0008_fd07920044cbf1b6fcf3603ebb166c6c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0007_13fbd5c85aa99a91119179870858fc36 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0007_d69fa1ea920d7f3f8644d49fc0b6aded interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0007_86034fd5bf9a19faed074c390256416d interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0007_8a643673abce113f9fea5ce1702250a6 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0007_265661e823b3aed3ebfb5ca65c6bba29 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0007_8b66ea616b172e9ee49fd7acdf30b938 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0007_dd667cce68b8575b61aaaa28533db684 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0007_5b4064ed641abac0e70005efb3a4ced2 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_5046950504d8a3b58a4241a8d69a199f interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_a78ba5b57311d51ec885c88fcf62a0d3 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0006_f476294ed42363d7c2356114aab61ba3 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_7a5fa0f486ff80bb9c7fb0e446b45447 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0006_6ebbdc3a0d909a8ad477a64e2c788e4f interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0006_db09fdc7c55f1435e80abe7fc49f0851 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0006_976d373277dceea1dca120722cddeb69 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0006_8c4a4d51c5e2a5272c519b0c73f39980 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0006_4139e5345038cc13f9683ec9a592592c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_e25c3a362b3347dd39e21d3465ef0049 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0006_34d5ca3511ee070149519075ad8a1d35 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_bbd5fd8057adff78016752e1be0a9d10 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_62f57ede9f4fdca4008bb8b102bff7ba interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0006_f696e01b39fbfe1dd45edc9adeb4154a interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0006_6bb429c97c0ecdbf2ad34c4ca1332c21 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_aa58b7d423443aa26503c4665dc3743d interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0006_4186142f3f5124cf0aa4c9fcf856b2c2 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0006_fdb81facd138c4a0705bd634d6c5ec20 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0006_ec5f3f3ec9a278b63dda59c8c26ba82e interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0006_5802b2f9aac2139509d07fe2b7e4c96b interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_b023136fe1c19ee81c3975c762b0b8e8 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0006_21c6aaf05ebd245d2a81458a7c024555 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0006_d079c9555b86aa5a077b492b67851ebc interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_59a3746398209af8d5abd1b9cc6e3ec8 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0006_8f10b08e9d138b1d7c3584d1ecc1451e interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0006_6f5a532ec7684dbd100cad131c0e6938 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0006_a0bfa772ad0992f8541f658be51b7492 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0006_8dc061cdbbf61be8784e703979185c75 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_5698322bcb20d1da317fbe9951ea1ea8 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_c670b45682d64d3a2f572ae08687b1c3 interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_d98ff0a96153b4e2dd5d28a88b553717 interface {
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0005_95dabc404203cf349885a9d084b8cd6d interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_036cfac2479cb17c204aeee52931f723 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_aa238e8cc38c595d62af9027950c4b6d interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0005_20631a232115a4d613bbdd90bd786e91 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_bb0588c994e9ade1989ee58ccd0341e0 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_2b9c34cec984ef83aa9d7eb139b39e1e interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_58c6e55019354e20f94dc97b5f01710e interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_020e8fd0be2f1cfda622cd7fae226094 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_c19e3acd7931b2ef3bd181d17c473d6c interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0005_13d3220ae21dcb3631a0e0d628adb44b interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0005_455bce9dcc99aa7c0fe96735f9c1c93c interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_172f2782b0ad97eecc7f56f17323847c interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0005_e5655e77364e3f931a3ea6875c01a0e7 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.SessionResetter
type wrapConn0005_26c5be906ea66f39030cc40db09bbbef interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_a6894db2a6de9012f66db439c1729127 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_86594d133db62770dc866e4b90a41484 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_be4447c4210552d0e2d709e04422e46c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_9820f7ab194cfc8681fa530ec810fbee interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_20d8d229b2484368b368c970e6deaaec interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_fd14a18f00f97dfb88beee48134b1ed7 interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_612cbcd42b1ca988c32b19b0ee60fe8e interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0005_858ad5997c7a1ecd0c4e016da0b78cb3 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0005_0ce7c7f80ec4cb8b2af6490440fb797b interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_6134bd04f1e6222f7503813d44eae9a4 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_522573917a875da9c8a8afd3f5bc8e45 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_c391b2ea017883e989dd4c91421a2518 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0005_afdced595d6aac26d33b2396172517d3 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_c4039cd70e55947c4038bbc29a3711d7 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_6f9178fa3ebe1aed2696357cd2afb3a9 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_650d3e7f112f289ca18b3e91964181c8 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_2e8be41ad679566886fbb5aa5471154b interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_1f8c595ee461d1900fce4cab06c9d1a3 interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_a03e3132f8a1d5217b2ffa017ef0e471 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_5e65f4368961d057065044067e0fdfa0 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_61a248bb5bae07d449dffffcab033bd3 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0005_48bc1b8fb59601a4f7543709bef7460b interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0005_8e89c1bd95e4e9731efc7ac1abba3639 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0005_370f70e1f7779d2b18c16b77796e6ad4 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Validator
type wrapConn0005_dccec1fa0b1ac6076c36fab915923e7c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0005_db318a810481bb79ec5c2f495823168c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext
type wrapConn0005_7a8098e8711012274339764352391092 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0005_0e4f2ce6ee8c6aaba65ea53cad0a7eda interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0005_dab2c71afcbd6af7c7e4e507d436ebb3 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.Validator
type wrapConn0005_8ea2641f961a5c26a90da3a80a453961 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0005_65e1c77325da4d0c054e69bd4551c599 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0005_c829a7b8ed999b05fe10801538e8fc57 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0005_3dde8f11372ab3df00c06e9cc9f532a0 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0005_436b3b6f6e649ed483a53237c3775c95 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.Validator
type wrapConn0005_497c507b429fdc6223283c0a0a534ea5 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0005_81570b25b723de4082fb023c40e4ccd9 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0005_063896be35045b544dfdc8bd6e379acb interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0005_646f19e9526b1b663e651ef0e2bef0d1 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0005_3c446fe7d7a33ac81540f3f0173a2617 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0004_a0bae6dc5e6fadde4228e11bf5e49dd1 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext
type wrapConn0004_28c597b75ba228e5d6902ddc68ad7811 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.QueryerContext
type wrapConn0004_d3598002e745fefcab218230adaf46d1 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ExecerContext|driver.QueryerContext|driver.Validator
type wrapConn0004_4ff0f4092dd79c548ab0e495e6cddc3b interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.Validator
type wrapConn0004_4ef020c4936aea427d194c2b96378df5 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0004_f018b29703e3acb35c78e692c0b4e037 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0004_0fd04b4e30706876c57056f2a95c5b01 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0004_c818cf1c81182b6fe0e9f65adf919fcc interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0004_c8f1b5b5a8743b962a62682d4a6e9262 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0004_15e77a0e9aeb76ab2f04aa748ac4719e interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0004_86042fa2173f60c1aaa263c2c4f6b7ee interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0004_645074a289d63811b9e9a8978d5ba260 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.Validator
type wrapConn0004_2a24acd339cf38eb1b877ff6e36477c2 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.Validator
type wrapConn0004_72837a807cd7df18bbd7e22ee697c816 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.Validator
}

// driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_e096a7508a321da18cfa19aba7f60b47 interface {
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.Validator
type wrapConn0004_62deeb4f7a0c751fdee63cefb834691f interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0004_573d41654139e877dea050658666ffa7 interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0004_84343f2e4ec9d07e1a9dbe14088de505 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0004_d3d3167450d7d79e6e587b43b5014c17 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0004_555cdb27b272a3f2c0eb9f7b9e0eb5ff interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Validator
type wrapConn0004_8fb2e6de52e01bddf84ba9aa01f20489 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Validator
type wrapConn0004_052895537635d2f5a24651bcc48ec975 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0004_f39e7f23bc4a7c1bd04d47bcdf9f2b75 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Validator
type wrapConn0004_8c1f1ab2627a320478eb2bd9c49efd73 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0004_d1052a967ff8db9233909609d294e184 interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0004_51825a78fc953cc98d20358550621612 interface {
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0004_95c3412bca03d21c7267d3279e2f3312 interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Validator
type wrapConn0004_684b88bf9d678a38c20afbfe31f4589a interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.SessionResetter
type wrapConn0004_089675b7ca4c9fa833b6515226a0b60c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0004_4f64f93a4e925dbd2a7598b37bb8f332 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0004_994bda6d7619bde46104d2eab7ff768c interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0004_bf56dd41a04eebaa9037e34d998af495 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ExecerContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_a602fc8de20f9667dca43dab54118301 interface {
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger|driver.SessionResetter
type wrapConn0004_8fef826c40cffbdc6a69c70b5676c824 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.QueryerContext|driver.Validator
type wrapConn0004_4501f160909b47a29559b718409520ea interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_87e9112616a76f4e1f0250177acda175 interface {
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.SessionResetter
type wrapConn0004_5d17fc1b7064520ae6022fd2f62c2bfb interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext
type wrapConn0004_65ab36fec2bb30a1596eb25d39a63d3c interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0004_ec4d0075fd0ddc8b390599c00aa63ac1 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.Pinger|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_ff5d54172f489033d905c69de9288f35 interface {
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.SessionResetter|driver.Validator
type wrapConn0004_4eafb813d5396d818b628d6737e80d42 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0004_1666fda37418e8810185898fd16573d2 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.NamedValueChecker|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_b214fbc2704c5d1b1313515138517b76 interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_86f21eab726445da5b1ba4da19a01e24 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_b454cff4daf9ab30e3870da405d5fc28 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0004_002ba8ae7ab8583e9fe9819c79f7f6e6 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ExecerContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_3b5220a9e11f03a530f5405a48a3583d interface {
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0004_8480b7a2f4b03ada42be3b2e1e8ee03b interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_421d78dfe11851b6277447d6fc3d762b interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_9efc0e3a9cf918d9d10523202afb880a interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0004_f33d7e93149b996e3b7d599a621eb41b interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0004_06ee4205d01722f7549f02e0b7087b0b interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0004_6afb89b2559180fa651920ad74abbb9c interface {
	driver.ConnBeginTx
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger
type wrapConn0004_5423a435d10b320ffac1aac27fe43c6f interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
}

// driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_457719d02ae8d6b7caabc81e5ac31d57 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0004_c44f3469765232bf2f3748ac3491ee68 interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_393cb9100ac7954d3e25afc8842b5dcb interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_09adbe5008d26d1ca9b6390889d3e56d interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0004_6e5b2720e60eff99a0177f056c13773b interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0004_c68aa300799a706219ed2f7ad15362c5 interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_1b5b9005d979312b723d3a27c0c65db6 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_7c2de3c593da1d0c0ca997bfb7e1d325 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0004_c6e51aefd95b2286637499cab5834392 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0004_25d1043092b1d4fde69e802678864cbe interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ExecerContext|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0004_ff336118365ff6fcd2a2e761486717c7 interface {
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0004_1d4652a872051fa3240ca534bef0a43e interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker
type wrapConn0004_b0e5c7fb09f6acbdfcf6e0d4258d4a9b interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
}

// driver.NamedValueChecker|driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0004_8098fe71a9a9219d58e590ca89192425 interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger|driver.SessionResetter
type wrapConn0004_6fcd36e5fdd1db57ef48d5bdd4beebd9 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger|driver.QueryerContext
type wrapConn0004_47c38651973c418d8726100d2bdf9d3a interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.QueryerContext|driver.SessionResetter
type wrapConn0003_23392d2c642abd59125f77ecca6f7b8f interface {
	driver.ConnBeginTx
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0003_67e361827a469e5a7fd41b0778aa1958 interface {
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ExecerContext|driver.QueryerContext|driver.SessionResetter
type wrapConn0003_c18fe73278d7622bc6a1b7ff02b795e7 interface {
	driver.ExecerContext
	driver.QueryerContext
	driver.SessionResetter
}

// driver.NamedValueChecker|driver.QueryerContext|driver.SessionResetter
type wrapConn0003_9d0152a9df319a9e7e4a0bdc7f7fa2d4 interface {
	driver.NamedValueChecker
	driver.QueryerContext
	driver.SessionResetter
}

// driver.NamedValueChecker|driver.Pinger|driver.SessionResetter
type wrapConn0003_75415c9d129d1dd60031d39ea83f5316 interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.SessionResetter
}

// driver.Pinger|driver.QueryerContext|driver.SessionResetter
type wrapConn0003_d40c325b5125e64636b9b8ea3f1b8975 interface {
	driver.Pinger
	driver.QueryerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.ExecerContext
type wrapConn0003_0980f76b2d444d923da26fc4a1b82818 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.ExecerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.NamedValueChecker
type wrapConn0003_1a34bf796737439096bfa6f10d5696c2 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.NamedValueChecker
}

// driver.ExecerContext|driver.Pinger|driver.SessionResetter
type wrapConn0003_3371fc70ba5d416a901c5118687e778e interface {
	driver.ExecerContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.NamedValueChecker
type wrapConn0003_ddd509d4dc401e51bac497647969675e interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.NamedValueChecker
}

// driver.ConnPrepareContext|driver.Pinger|driver.SessionResetter
type wrapConn0003_66e6b00152e5329959d8d66776407655 interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.Pinger|driver.SessionResetter
type wrapConn0003_ce39b6405b9206d7ba9fe332b19cf89f interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.NamedValueChecker
type wrapConn0003_ff274753683e1d0adf6ecc55b9be8631 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.NamedValueChecker
}

// driver.QueryerContext|driver.SessionResetter|driver.Validator
type wrapConn0003_613592023fc524d8f895420661ca88e0 interface {
	driver.QueryerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Validator
type wrapConn0003_ecd9017910d8eaafad4c84916460cde2 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Validator
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.Pinger
type wrapConn0003_4afcf6b296cf46480b7f20345361eec9 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.Pinger
}

// driver.ExecerContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0003_2c310269b09748f7be39d17b8528f2df interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Validator
type wrapConn0003_214d339ea349168fbe3e0955424c5eef interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0003_4b5d9a137b792a784a0b494e06773edb interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.SessionResetter
type wrapConn0003_67a6f30839ce48a2e9f3488bf5e2053d interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.Pinger
type wrapConn0003_1bf343887515a244e3d82298438693bd interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.Pinger
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Validator
type wrapConn0003_3046c3f6840014e9d6851479421aa2db interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.SessionResetter
type wrapConn0003_429af3d54eff18d5961d7d87b53413b7 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.ExecerContext|driver.SessionResetter
type wrapConn0003_348becb0e4f2e9da375f40a1e59807d1 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.Pinger
type wrapConn0003_1c8897aa6698e41ba80853991cc24839 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.Pinger
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.SessionResetter
type wrapConn0003_c02e55bb6f1875e5ca0c1667579be31a interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.SessionResetter
}

// driver.Pinger|driver.SessionResetter|driver.Validator
type wrapConn0003_d033831362fd7cd8abd09076a328bf5f interface {
	driver.Pinger
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Pinger
type wrapConn0003_5a07d33df77df1576ce8067b8997d413 interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ConnPrepareContext|driver.ExecerContext|driver.QueryerContext
type wrapConn0003_58164b71668ef615a11f9d6dd3ff1e07 interface {
	driver.ConnPrepareContext
	driver.ExecerContext
	driver.QueryerContext
}

// driver.NamedValueChecker|driver.SessionResetter|driver.Validator
type wrapConn0003_f3e446fd938f183e46a975476439f43d interface {
	driver.NamedValueChecker
	driver.SessionResetter
	driver.Validator
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.Validator
type wrapConn0003_ddca4fe2167f6deb65f9ee0d4d46b23b interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Validator
type wrapConn0003_55276ca2338c0cd130c39fb53ac4caaa interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Validator
type wrapConn0003_344b3b12f6cdf795bbc480ca3c22d0eb interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0003_68f463fe3d9a985f2c57f5f872c7c333 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ConnBeginTx|driver.Pinger|driver.Validator
type wrapConn0003_b2d88477280fe09a57ad676180ae7fca interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger|driver.Validator
type wrapConn0003_aebab2fa61e210e3c96cb924b4356379 interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.Validator
}

// driver.NamedValueChecker|driver.Pinger|driver.QueryerContext
type wrapConn0003_55b0c9c15fb89c88e545e8b9afe973ee interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.QueryerContext
}

// driver.ExecerContext|driver.Pinger|driver.Validator
type wrapConn0003_57e946d049a38df9bcfe6ef5d8598cfc interface {
	driver.ExecerContext
	driver.Pinger
	driver.Validator
}

// driver.NamedValueChecker|driver.Pinger|driver.Validator
type wrapConn0003_4e478d169a0c925294ee737c49c1ac6a interface {
	driver.NamedValueChecker
	driver.Pinger
	driver.Validator
}

// driver.ExecerContext|driver.SessionResetter|driver.Validator
type wrapConn0003_065f30b751342fccddaa383211967ee5 interface {
	driver.ExecerContext
	driver.SessionResetter
	driver.Validator
}

// driver.ExecerContext|driver.Pinger|driver.QueryerContext
type wrapConn0003_80b435a335887493ba9656f5a0e76a14 interface {
	driver.ExecerContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.QueryerContext|driver.Validator
type wrapConn0003_975ad90f36f2a8d5d14cbc769a6ce607 interface {
	driver.ConnBeginTx
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger|driver.QueryerContext
type wrapConn0003_ecd209322f30837e1dde5cc3c5903232 interface {
	driver.ConnPrepareContext
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.Pinger|driver.QueryerContext
type wrapConn0003_f589cedd98232e73dba33f2b1ee57f05 interface {
	driver.ConnBeginTx
	driver.Pinger
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.SessionResetter|driver.Validator
type wrapConn0003_895867613f93da07976594c2ea883249 interface {
	driver.ConnPrepareContext
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.QueryerContext|driver.Validator
type wrapConn0003_6168ba655398488893634accebfceca9 interface {
	driver.ConnPrepareContext
	driver.QueryerContext
	driver.Validator
}

// driver.ConnBeginTx|driver.SessionResetter|driver.Validator
type wrapConn0003_7908fe1d05a5c3552e8ea5c1ee7d48e7 interface {
	driver.ConnBeginTx
	driver.SessionResetter
	driver.Validator
}

// driver.NamedValueChecker|driver.QueryerContext|driver.Validator
type wrapConn0003_0040353f748f8b98cb8da172c6b336c1 interface {
	driver.NamedValueChecker
	driver.QueryerContext
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0003_42faf5a4eb464290a69fa9800a551bd4 interface {
	driver.ExecerContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.Pinger|driver.QueryerContext|driver.Validator
type wrapConn0003_ee7a3b8aa7f629ea0410d3cb5df23364 interface {
	driver.Pinger
	driver.QueryerContext
	driver.Validator
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0003_4a1ee3cec4754da65ab631f8d55709d8 interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.NamedValueChecker|driver.QueryerContext
type wrapConn0003_86f0c8258522256415a6791126d2786c interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext|driver.QueryerContext
type wrapConn0003_cdbed2933e5c0be5f6806d38207387b1 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ExecerContext|driver.QueryerContext
type wrapConn0003_66ae780e5764601ed26339f0ef146ab9 interface {
	driver.ConnBeginTx
	driver.ExecerContext
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.NamedValueChecker|driver.Pinger
type wrapConn0003_e8d66097d62d92ca5e8a01d0627ef47b interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ExecerContext|driver.QueryerContext|driver.Validator
type wrapConn0003_086984675e3daca69cb4d93a08b03c05 interface {
	driver.ExecerContext
	driver.QueryerContext
	driver.Validator
}

// driver.ExecerContext|driver.QueryerContext
type wrapConn0002_9998c1796991ed1018aa6832ef10c20d interface {
	driver.ExecerContext
	driver.QueryerContext
}

// driver.NamedValueChecker|driver.QueryerContext
type wrapConn0002_608b01490fc5eeb54196421e44a1d294 interface {
	driver.NamedValueChecker
	driver.QueryerContext
}

// driver.ConnPrepareContext|driver.QueryerContext
type wrapConn0002_e5960e40a7bbe1b5a19c38f71e0e6861 interface {
	driver.ConnPrepareContext
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.QueryerContext
type wrapConn0002_21490f92aea80b01cad7a6a9b496de9d interface {
	driver.ConnBeginTx
	driver.QueryerContext
}

// driver.ConnBeginTx|driver.ConnPrepareContext
type wrapConn0002_b7ff310fb7e22a7ccbb847bf26429dd9 interface {
	driver.ConnBeginTx
	driver.ConnPrepareContext
}

// driver.QueryerContext|driver.SessionResetter
type wrapConn0002_734bbe2891841d705d0385cb4817474b interface {
	driver.QueryerContext
	driver.SessionResetter
}

// driver.Pinger|driver.QueryerContext
type wrapConn0002_fb43968160809b260946ae6f55b7fe56 interface {
	driver.Pinger
	driver.QueryerContext
}

// driver.QueryerContext|driver.Validator
type wrapConn0002_57475ba369865550e7711feefe48c5d4 interface {
	driver.QueryerContext
	driver.Validator
}

// driver.Pinger|driver.Validator
type wrapConn0002_dea76879684fe95f8ca4bec7f7cec2d7 interface {
	driver.Pinger
	driver.Validator
}

// driver.NamedValueChecker|driver.Validator
type wrapConn0002_9743a8d2379220bf3bef64c2a633bd0d interface {
	driver.NamedValueChecker
	driver.Validator
}

// driver.ConnBeginTx|driver.ExecerContext
type wrapConn0002_01ce0f63e6a6b920df90cbea6b6ede03 interface {
	driver.ConnBeginTx
	driver.ExecerContext
}

// driver.ConnBeginTx|driver.SessionResetter
type wrapConn0002_f21346c5fd58f22896f02b5b5ed3475e interface {
	driver.ConnBeginTx
	driver.SessionResetter
}

// driver.NamedValueChecker|driver.Pinger
type wrapConn0002_0921fb5e4c916e10370d5e0398435d4e interface {
	driver.NamedValueChecker
	driver.Pinger
}

// driver.ConnPrepareContext|driver.SessionResetter
type wrapConn0002_1bf71fb4be83d1fefe01cdd4d5a7a547 interface {
	driver.ConnPrepareContext
	driver.SessionResetter
}

// driver.ConnPrepareContext|driver.NamedValueChecker
type wrapConn0002_35700b3f7c770fcd0bf83836eee99fec interface {
	driver.ConnPrepareContext
	driver.NamedValueChecker
}

// driver.NamedValueChecker|driver.SessionResetter
type wrapConn0002_c4cbf403989171fe81ee42f1529c60fe interface {
	driver.NamedValueChecker
	driver.SessionResetter
}

// driver.ExecerContext|driver.Pinger
type wrapConn0002_b0936442b3763765deb78c0f8945bb0a interface {
	driver.ExecerContext
	driver.Pinger
}

// driver.SessionResetter|driver.Validator
type wrapConn0002_20cb62afffd12ef5a3dd67710522c25c interface {
	driver.SessionResetter
	driver.Validator
}

// driver.ConnPrepareContext|driver.Pinger
type wrapConn0002_039f074dd8b7cd2fd02bcb0f3a89e29a interface {
	driver.ConnPrepareContext
	driver.Pinger
}

// driver.ConnBeginTx|driver.Pinger
type wrapConn0002_defc836eab973764559fef7a6eacab29 interface {
	driver.ConnBeginTx
	driver.Pinger
}

// driver.ConnPrepareContext|driver.ExecerContext
type wrapConn0002_f2e0c829a13f0fc5e34c7487f6a5743b interface {
	driver.ConnPrepareContext
	driver.ExecerContext
}

// driver.ConnPrepareContext|driver.Validator
type wrapConn0002_739d949041a357f14dd543ae1c82d4b4 interface {
	driver.ConnPrepareContext
	driver.Validator
}

// driver.Pinger|driver.SessionResetter
type wrapConn0002_004b3f5abaa12288a30861da40988df6 interface {
	driver.Pinger
	driver.SessionResetter
}

// driver.ConnBeginTx|driver.Validator
type wrapConn0002_7eb98af2c595d337fef40bcf00e367f5 interface {
	driver.ConnBeginTx
	driver.Validator
}

// driver.ExecerContext|driver.NamedValueChecker
type wrapConn0002_70a606d71e13e67187acc5d5056da2e1 interface {
	driver.ExecerContext
	driver.NamedValueChecker
}

// driver.ConnBeginTx|driver.NamedValueChecker
type wrapConn0002_1c65fb5735328219b65d13d4e717b79f interface {
	driver.ConnBeginTx
	driver.NamedValueChecker
}

// driver.ExecerContext|driver.SessionResetter
type wrapConn0002_26864a2d0c8c5a605e497ec245ca9aa1 interface {
	driver.ExecerContext
	driver.SessionResetter
}

// driver.ExecerContext|driver.Validator
type wrapConn0002_b1c510487a94a905aff86496f31078e6 interface {
	driver.ExecerContext
	driver.Validator
}

// driver.NamedValueChecker
type wrapConn0001_dd0fe84666fad76cbf0dee440e953d06 interface {
	driver.NamedValueChecker
}

// driver.ConnBeginTx
type wrapConn0001_d815927e8e5ae23caed6edbf82ddb9a8 interface {
	driver.ConnBeginTx
}

// driver.Pinger
type wrapConn0001_f95e29e11298b900f8a05d72dcfe041b interface {
	driver.Pinger
}

// driver.SessionResetter
type wrapConn0001_70145ca1951f1697e67aa095f55b8305 interface {
	driver.SessionResetter
}

// driver.ExecerContext
type wrapConn0001_f98304698b6793d84b0e7be539d135ce interface {
	driver.ExecerContext
}

// driver.QueryerContext
type wrapConn0001_7a610c95e7c563485e18c228cc42e7fb interface {
	driver.QueryerContext
}

// driver.ConnPrepareContext
type wrapConn0001_c3128ba459e963c7f07c0de8d58b775f interface {
	driver.ConnPrepareContext
}

// driver.Validator
type wrapConn0001_abc062388a1f7a907b94470050b8bc86 interface {
	driver.Validator
}

func wrapStmt(stmt driver.Stmt, query string, opts Options) driver.Stmt {
	c := &wrapperStmt{stmt: stmt, query: query, opts: opts}
	if _, ok := stmt.(wrapStmt0003_0856d47f81a59090013bce4a1ee5bdd5); ok {
	return struct {
		driver.Stmt
			driver.StmtExecContext
			driver.StmtQueryContext
			driver.NamedValueChecker		
}{c, c, c, c}
}

	if _, ok := stmt.(wrapStmt0002_90dca7297ea2f892831234b29884308e); ok {
	return struct {
		driver.Stmt
			driver.StmtQueryContext
			driver.NamedValueChecker		
}{c, c, c}
}

	if _, ok := stmt.(wrapStmt0002_21dd080dbc547db7e0f1b0faf56aa2c3); ok {
	return struct {
		driver.Stmt
			driver.StmtExecContext
			driver.NamedValueChecker		
}{c, c, c}
}

	if _, ok := stmt.(wrapStmt0002_1ac318a0d0640d654e838a3d210e928f); ok {
	return struct {
		driver.Stmt
			driver.StmtExecContext
			driver.StmtQueryContext		
}{c, c, c}
}

	if _, ok := stmt.(wrapStmt0001_dd0fe84666fad76cbf0dee440e953d06); ok {
	return struct {
		driver.Stmt
			driver.NamedValueChecker		
}{c, c}
}

	if _, ok := stmt.(wrapStmt0001_3f3eb31555c2ec79768169cb473a84fe); ok {
	return struct {
		driver.Stmt
			driver.StmtQueryContext		
}{c, c}
}

	if _, ok := stmt.(wrapStmt0001_be979ee6572ef74e683e3153d6b71c90); ok {
	return struct {
		driver.Stmt
			driver.StmtExecContext		
}{c, c}
}

return c
}
// driver.StmtExecContext|driver.StmtQueryContext|driver.NamedValueChecker
type wrapStmt0003_0856d47f81a59090013bce4a1ee5bdd5 interface {
	driver.StmtExecContext
	driver.StmtQueryContext
	driver.NamedValueChecker
}

// driver.StmtQueryContext|driver.NamedValueChecker
type wrapStmt0002_90dca7297ea2f892831234b29884308e interface {
	driver.StmtQueryContext
	driver.NamedValueChecker
}

// driver.StmtExecContext|driver.NamedValueChecker
type wrapStmt0002_21dd080dbc547db7e0f1b0faf56aa2c3 interface {
	driver.StmtExecContext
	driver.NamedValueChecker
}

// driver.StmtExecContext|driver.StmtQueryContext
type wrapStmt0002_1ac318a0d0640d654e838a3d210e928f interface {
	driver.StmtExecContext
	driver.StmtQueryContext
}

// driver.NamedValueChecker
type wrapStmt0001_dd0fe84666fad76cbf0dee440e953d06 interface {
	driver.NamedValueChecker
}

// driver.StmtQueryContext
type wrapStmt0001_3f3eb31555c2ec79768169cb473a84fe interface {
	driver.StmtQueryContext
}

// driver.StmtExecContext
type wrapStmt0001_be979ee6572ef74e683e3153d6b71c90 interface {
	driver.StmtExecContext
}


