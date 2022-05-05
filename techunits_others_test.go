package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-examples-client/techunits"
)

var conn *grpc.ClientConn
var c pb.SampleServiceClient

func setup() {
	// Set up a connection to the server.
	var err error
	conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = pb.NewSampleServiceClient(conn)
}

func teardown() {
	conn.Close()
}

func setupSubtest(t *testing.T) {
	setup()
	t.Logf("[SETUP] connection")
}

func teardownSubtest(t *testing.T) {
	teardown()
	t.Logf("[TEARDOWN] Bye!")
}

func setupTest(m *testing.M) {
	setup()
}

func teardownTest(m *testing.M) {
	teardown()
}

func TestMain(m *testing.M) {
	setupTest(m)
	defer teardownTest(m)
	os.Exit(m.Run())
}

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		res  string
		err  error
	}{
		{
			"default",
			pingResponse,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			response, err := c.Ping(ctx, &pb.PingRequest{})

			if err != nil {
				t.Error("error: received", err)
			}

			response.Id, response.CreatedOn = "", 0
			b, _ := json.Marshal(response)
			s := string(b)

			if s != tt.res {
				t.Error("error: expected", tt.res, "\n\nreceived", s)
			}
		})
	}
}

func TestEntryListRequest_2nd(t *testing.T) {
	tests := []struct {
		name string
		res  string
		err  error
	}{
		{
			"default",
			entryList3,
			nil,
		},
	}

	// setupSubtest(t)
	// defer teardownSubtest(t)
	//c := pb.NewSampleServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Contact the server
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			response, err := c.ListEntries(ctx, &pb.EntryListRequest{})

			if err != nil {
				t.Error("error: received", err)
			}

			b, _ := json.Marshal(response)
			s := string(b)

			if s != tt.res {
				t.Error("error: expected", tt.res, "\n\nreceived", s)
			}
		})
	}
}

const pingResponse = `{"field_01":"N/A","field_02":"N/A","field_03":"N/A","field_04":"N/A","status":"OK"}`

const entryList2 = `{"entries":[{"id":"0309694d-0888-42a6-9577-ad19bf0a06e4","title":"explicabo explicabo harum","code":"explicabo-explicabo-harum","description":"Odit libero quia neque rerum velit. Esse porro fuga. Quis fugiat nam.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/0309694d-0888-42a6-9577-ad19bf0a06e4/a949950e-8c86-47b7-98b2-8aa38613434a.png","created_on":1651779171,"modified_on":1651779171},{"id":"6e53ebc7-d3f9-420f-ab1f-e5e0e117b6fc","title":"molestias quo minus","code":"molestias-quo-minus","description":"Dolores rerum error maxime vero vel facilis. Eius necessitatibus blanditiis quia qui voluptatibus voluptatem et et aut. Repudiandae voluptatum magnam porro fugiat ducimus officia.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/6e53ebc7-d3f9-420f-ab1f-e5e0e117b6fc/87a57b37-afe2-4e69-b8b9-bb1a71a43a4f.png","created_on":1651779171,"modified_on":1651779171},{"id":"b84fca93-ef73-409b-9a18-fae825ea8a9e","title":"et rerum ad","code":"et-rerum-ad","description":"Nihil aut provident enim ut et ratione at cum. Corrupti ullam et voluptatem. Nulla et itaque a velit id in iste. Ex ad beatae voluptas omnis numquam sed quidem ut. Deserunt doloribus voluptate qui tenetur.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/b84fca93-ef73-409b-9a18-fae825ea8a9e/a40bd68d-4922-4792-91e9-df319ebae8b4.png","created_on":1651779171,"modified_on":1651779171},{"id":"960becf5-6d27-4f8d-908c-5cd39179e94f","title":"modi asperiores nulla","code":"modi-asperiores-nulla","description":"Fuga distinctio quam autem aliquam. Et sed exercitationem iste nihil odit. Pariatur itaque quo ipsa modi. Consectetur omnis omnis aspernatur aut nihil alias.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/960becf5-6d27-4f8d-908c-5cd39179e94f/a9e2349f-572e-46aa-aadc-ce51af3ab7c4.png","created_on":1651779171,"modified_on":1651779171},{"id":"98e20eb4-77a7-4ade-9d28-f753613f9d1b","title":"quae cumque voluptatum","code":"quae-cumque-voluptatum","description":"Mollitia nihil eos qui ducimus et eaque consequatur eos. Autem alias nesciunt nobis velit aliquam aut eos ea natus. Qui dignissimos et perferendis cum consequatur non dolorem adipisci. Et voluptatibus voluptatem tempore qui dolorem voluptas dolorem nemo. Autem asperiores perferendis magnam beatae et.","created_on":1651779171,"modified_on":1651779171},{"id":"3f33e089-19ce-4da9-a5b6-e24501a487a8","title":"iste qui atque","code":"iste-qui-atque","description":"Fugiat hic in deleniti consectetur itaque. Consequuntur nesciunt sunt. Rerum aspernatur explicabo occaecati odit tempora reprehenderit voluptas.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/3f33e089-19ce-4da9-a5b6-e24501a487a8/ce9e924a-3762-4719-84fe-61821a25852c.png","created_on":1651779171,"modified_on":1651779171},{"id":"3c20c24f-5b14-42ac-bf47-29473e709757","title":"quidem odio commodi","code":"quidem-odio-commodi","description":"Consequuntur aut sit quos voluptatem est nihil sint nobis ea. Dolorum vel eos fugiat molestiae necessitatibus. Numquam sit quod modi dolorem sed explicabo aut aliquid recusandae. Et adipisci soluta rerum dolores repellendus voluptatem unde molestiae consequatur. Ea eos in esse qui itaque.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/3c20c24f-5b14-42ac-bf47-29473e709757/80d10d1a-ee4c-443f-a212-0177f661a300.png","created_on":1651779171,"modified_on":1651779171},{"id":"80e6467a-ccc1-4e47-8de3-f0c71949f552","title":"nam excepturi eligendi","code":"nam-excepturi-eligendi","description":"Architecto soluta ad repudiandae nemo inventore excepturi temporibus. Voluptas aut cumque. Impedit explicabo non similique deserunt totam tempora assumenda.","created_on":1651779171,"modified_on":1651779171},{"id":"cf9eac38-c622-4a45-bafc-709f3a15e585","title":"alias iure et","code":"alias-iure-et","description":"Quia corrupti amet odit explicabo quaerat minima aliquam et enim. Beatae vitae doloribus quis illum et ut. Debitis ipsa aut soluta fuga deserunt vel magnam. Esse reiciendis fugit est id modi.","created_on":1651779171,"modified_on":1651779171},{"id":"6f4890c5-9f2d-43b7-9061-397b0d632510","title":"et nulla accusamus","code":"et-nulla-accusamus","description":"Voluptatem eaque voluptatem unde sit ducimus repellendus maxime est inventore. Mollitia maiores ea et delectus nostrum ipsum omnis numquam vel. Quo molestiae voluptas corrupti.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/6f4890c5-9f2d-43b7-9061-397b0d632510/3cc5add5-90a0-41dc-a0a2-9f3483a6d213.png","created_on":1651779171,"modified_on":1651779171},{"id":"b2822258-fef0-49f8-b67c-23d5d347f0fd","title":"aliquid et nemo","code":"aliquid-et-nemo","description":"Accusamus aut nihil voluptatem fugiat ex est. Delectus repellat et quasi. Vel sunt aliquid molestiae aut earum qui eum placeat. Dolores non saepe sed. In et tempore excepturi consequuntur dolor non nesciunt.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/b2822258-fef0-49f8-b67c-23d5d347f0fd/62a2311b-347f-4952-becb-78a589840f61.png","created_on":1651779171,"modified_on":1651779171},{"id":"6a67713a-438b-4824-86ad-23c0f238c0bf","title":"architecto ad hic","code":"architecto-ad-hic","description":"Ut ipsam incidunt. Atque ducimus esse ab aliquid. Nisi rem adipisci perspiciatis.","created_on":1651779171,"modified_on":1651779171},{"id":"3812c711-9af9-419a-b8da-4007adc6bf42","title":"eius ad voluptatibus","code":"eius-ad-voluptatibus","description":"Maiores et rerum id natus. Asperiores error sunt nemo quis harum quod quisquam dignissimos commodi. Itaque perspiciatis autem. Quia ut aut reiciendis quia.","created_on":1651779171,"modified_on":1651779171},{"id":"686064e9-ccb7-4438-9580-957fe04c616e","title":"et perspiciatis et","code":"et-perspiciatis-et","description":"Quod itaque rerum eaque sint. Cum et consequatur quo rerum odit praesentium optio. Consequatur quia nostrum quisquam repellendus tempora voluptates quo eius. Id ducimus excepturi quisquam rerum dolores at. Vel doloremque tempore. Sint inventore nisi illum asperiores eaque nulla.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/686064e9-ccb7-4438-9580-957fe04c616e/88c4aa1c-f148-4597-8b45-78bdf34329d9.png","created_on":1651779171,"modified_on":1651779171},{"id":"83d50077-b23f-449f-9068-b4fb79e59535","title":"est asperiores et","code":"est-asperiores-et","description":"Voluptate sequi ab. Harum rem molestiae nam voluptates. Corrupti deleniti facilis amet voluptas. Dolorem unde tempora ab.","created_on":1651779171,"modified_on":1651779171},{"id":"80eb895b-24f8-4242-b7f8-a941d70768f4","title":"magni voluptatem saepe","code":"magni-voluptatem-saepe","description":"Cupiditate molestias et sunt. Blanditiis nihil eum rerum eos harum aut qui. Consequatur suscipit qui exercitationem veniam illum optio facilis eum doloribus. Voluptatem aliquam omnis. Illo neque natus ad quo porro et commodi. Est eos quia laudantium vitae.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/80eb895b-24f8-4242-b7f8-a941d70768f4/20b3fe20-27ca-41dc-ab55-ac1110f0626d.png","created_on":1651779171,"modified_on":1651779171},{"id":"a152d93b-9183-4a70-9a73-838dd6900868","title":"nemo aut aliquam","code":"nemo-aut-aliquam","description":"Quibusdam ut minus ipsam iure vero itaque ipsam. Exercitationem qui consequatur. Optio aut unde.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/a152d93b-9183-4a70-9a73-838dd6900868/334fe4d1-6e8d-401e-86b4-62da388b6ac8.png","created_on":1651779171,"modified_on":1651779171},{"id":"0b48a0fc-4735-460c-bbef-ab14d9a6f79b","title":"veniam eaque porro","code":"veniam-eaque-porro","description":"Autem et dolore doloremque et temporibus. Earum cum enim quia magni. Quaerat exercitationem officiis. Pariatur distinctio distinctio. Beatae rerum nihil ea sequi eius laudantium voluptatem laborum. Sequi et cumque ut qui.","created_on":1651779171,"modified_on":1651779171},{"id":"a22e40bb-486d-4639-a94a-fa056814872f","title":"libero corrupti libero","code":"libero-corrupti-libero","description":"Odit asperiores placeat illo. Est ut laudantium. Necessitatibus ex nam qui sequi fuga praesentium maiores.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/a22e40bb-486d-4639-a94a-fa056814872f/db2a2c08-bdba-4138-ada0-35eb81fa154c.png","created_on":1651779171,"modified_on":1651779171},{"id":"5093269c-fd91-4f78-ba68-fa2b2a72a6e6","title":"culpa ab fugiat","code":"culpa-ab-fugiat","description":"Possimus facere quasi assumenda exercitationem illo ut recusandae animi. In commodi eaque quas harum. Et necessitatibus doloribus laborum ipsa occaecati vel quasi. Qui pariatur occaecati expedita minus. Error id suscipit temporibus.","created_on":1651779171,"modified_on":1651779171},{"id":"c4c91d0e-de61-4f83-98c4-356f5c200065","title":"dolores numquam in","code":"dolores-numquam-in","description":"Labore mollitia beatae ut. Non quam iusto amet. Quam adipisci deleniti labore. Doloribus illo omnis nihil optio cum perferendis iure.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/c4c91d0e-de61-4f83-98c4-356f5c200065/3591a3e3-aa99-4000-8f03-e1ecbbffcd42.png","created_on":1651779171,"modified_on":1651779171},{"id":"b0068bbd-5dc8-4def-bcb5-b2e79eed9566","title":"cupiditate autem dignissimos","code":"cupiditate-autem-dignissimos","description":"Deserunt officiis ea aliquam nesciunt nostrum et ut non. Iure reiciendis natus aut blanditiis ipsa aliquid incidunt reprehenderit. Voluptate in amet et omnis voluptas.","created_on":1651779171,"modified_on":1651779171},{"id":"9edf1beb-333a-41f9-91df-8011b437e5e0","title":"officiis velit ratione","code":"officiis-velit-ratione","description":"Blanditiis cumque accusantium quaerat adipisci qui totam non veritatis dicta. Alias earum porro vero. Optio rerum sint nihil quia suscipit vel. Et assumenda porro minima voluptatem beatae officiis. Sit est error hic ipsam.","created_on":1651779171,"modified_on":1651779171},{"id":"26718c4f-e454-45ec-9217-e6c59bfe00f3","title":"non nisi ex","code":"non-nisi-ex","description":"Voluptatibus et perferendis quia. Voluptas iste officia quo placeat doloribus consectetur. Aut quod explicabo molestias accusamus. Porro et quos qui ut tenetur beatae fugit. Possimus cupiditate repudiandae id sed exercitationem autem accusantium.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/26718c4f-e454-45ec-9217-e6c59bfe00f3/35e089fe-6008-47d5-a4b0-e21bb91a86a6.png","created_on":1651779171,"modified_on":1651779171},{"id":"8f194f15-f60b-4994-818f-c2e10eba7de6","title":"omnis ullam qui","code":"omnis-ullam-qui","description":"Distinctio eveniet dicta recusandae totam ipsa est. Temporibus voluptas dolorem odit blanditiis et expedita dolorem. Ad occaecati necessitatibus. Sunt sequi quibusdam sunt adipisci. Excepturi quia impedit consequuntur dolore est officiis omnis ipsa dolores.","created_on":1651779171,"modified_on":1651779171}]}`
const entryList3 = `{"entries":[{"id":"5983fea7-c3b7-4ac5-a469-9326992d5c90","title":"quis ipsam atque","code":"quis-ipsam-atque","description":"Ea saepe asperiores delectus voluptatum aut laboriosam. Sequi minima ad cupiditate ad nisi quidem. Non molestiae fugit.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/5983fea7-c3b7-4ac5-a469-9326992d5c90/6cfc7e87-d316-40eb-9df2-fcb5b2efbcc5.png","created_on":1651783674,"modified_on":1651783674},{"id":"132eb599-66e7-4bc1-ad97-64026c0e1ffa","title":"omnis est dolores","code":"omnis-est-dolores","description":"Itaque occaecati libero. Quisquam occaecati omnis aut omnis et omnis. Sint voluptate culpa.","created_on":1651783674,"modified_on":1651783674},{"id":"f60475b2-caac-44b0-b5d2-f47e23981c72","title":"adipisci doloribus aut","code":"adipisci-doloribus-aut","description":"Commodi unde velit. Ut vel temporibus. Velit dolores laborum est aut et nemo qui dolorum.","created_on":1651783674,"modified_on":1651783674},{"id":"7bb00fb9-678b-41d0-9bfc-6cd9cce51006","title":"incidunt aliquam iusto","code":"incidunt-aliquam-iusto","description":"Totam ullam et minima illo veniam quasi ipsa. Laboriosam temporibus ut corrupti voluptatum. Ducimus voluptates beatae nisi.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/7bb00fb9-678b-41d0-9bfc-6cd9cce51006/2e367cd5-6714-4847-a36c-e203652049ad.png","created_on":1651783674,"modified_on":1651783674},{"id":"533d14f9-1fe0-43d8-9a69-6e6393b7305b","title":"sit est sunt","code":"sit-est-sunt","description":"At quo maxime eos enim velit quis blanditiis explicabo. Ipsam enim quas ipsam. Earum a corporis ea laboriosam aperiam. Et nobis veritatis error adipisci quia similique qui. Mollitia quisquam reiciendis voluptas enim est aliquam dolorem laudantium. Autem reprehenderit sequi veniam.","created_on":1651783674,"modified_on":1651783674},{"id":"20038a2e-9760-403b-9399-13edbdc804a4","title":"dicta commodi cupiditate","code":"dicta-commodi-cupiditate","description":"Sed quidem nam beatae veniam quod. Fuga iure dolor est consectetur placeat. Possimus ipsum magnam enim dolor. Est facere dolor dolores enim sit dolores molestiae laborum.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/20038a2e-9760-403b-9399-13edbdc804a4/bbd1cb1c-adc3-4088-9590-db26754e57bf.png","created_on":1651783674,"modified_on":1651783674},{"id":"46144134-5dbf-41a7-93bd-c7a47700fd92","title":"voluptas pariatur maxime","code":"voluptas-pariatur-maxime","description":"Consequatur sapiente nulla sed aut qui dolores ea inventore sint. Ipsam qui eaque molestias. Distinctio quis unde nesciunt suscipit. Exercitationem autem qui.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/46144134-5dbf-41a7-93bd-c7a47700fd92/c1249ddf-c21a-41fe-baee-acd58779c781.png","created_on":1651783673,"modified_on":1651783673},{"id":"0f4c6144-2e61-41bd-9c2d-5ab516afeedb","title":"et velit et","code":"et-velit-et","description":"Expedita officiis deleniti id fugiat quaerat nihil necessitatibus ratione. Deserunt iste quam. Maxime est tenetur cum est. Odio quo illo consequatur quae voluptates sit ex. Consequatur sint cumque et ut tempora ullam.","created_on":1651783673,"modified_on":1651783673},{"id":"e9d86346-5083-46e8-8908-e34afcdc1355","title":"ipsum nobis aliquam","code":"ipsum-nobis-aliquam","description":"Doloribus ut omnis unde aspernatur beatae velit. Numquam quo eveniet est quam consequatur sequi aspernatur fugiat. Dolorum unde sequi non.","created_on":1651783673,"modified_on":1651783673},{"id":"d28fca9f-db53-4a20-92d7-c3c864cf898c","title":"aut earum doloribus","code":"aut-earum-doloribus","description":"Atque nihil et animi excepturi voluptates quasi deserunt. Voluptates facere doloribus et sit aut voluptates minima similique. Dignissimos et dicta sint non ea voluptate repudiandae earum.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/d28fca9f-db53-4a20-92d7-c3c864cf898c/3c0e4313-74fe-4e9a-8400-e4976eae0551.png","created_on":1651783673,"modified_on":1651783673},{"id":"cbc8e8b7-d59e-4601-b1ee-5dd7f8b6b030","title":"praesentium totam repellendus","code":"praesentium-totam-repellendus","description":"Quis tempora itaque architecto. Enim reprehenderit cum. Nobis odit omnis dolores ut voluptas aut qui et. Consequatur itaque nemo ea iste accusamus ab aut. Nesciunt accusantium quod ea error est unde sequi sit. Eos quo et quisquam quasi cupiditate quod qui.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/cbc8e8b7-d59e-4601-b1ee-5dd7f8b6b030/bc4db916-354d-4a9f-bac7-9e70c62488db.png","created_on":1651783673,"modified_on":1651783673},{"id":"db7dba2d-0cbd-4b8d-bb26-1579d1734092","title":"qui laboriosam odio","code":"qui-laboriosam-odio","description":"Rerum nulla tempora enim harum dolorum voluptas. Temporibus accusamus voluptate ut quod harum fuga. Dicta quo aut ipsam id incidunt aut quasi soluta magni. Enim illo ducimus. Autem inventore consequatur.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/db7dba2d-0cbd-4b8d-bb26-1579d1734092/16f4bac4-a47e-40dc-8ba4-3e49310e4574.png","created_on":1651783673,"modified_on":1651783673},{"id":"7ff7a0e1-64ad-49a1-93f2-077318660f74","title":"laboriosam maxime est","code":"laboriosam-maxime-est","description":"Officia et aut et dolorem soluta vitae excepturi alias consequatur. Cum qui ut natus et quo minima. Explicabo a sint occaecati accusamus pariatur. Qui officia maxime voluptas vitae ad.","created_on":1651783673,"modified_on":1651783673},{"id":"2b5b7288-0999-4f21-a1be-4bcf756ac1d1","title":"atque vero fugit","code":"atque-vero-fugit","description":"Mollitia eveniet et odit dolor reprehenderit aut. Fugiat aut est ullam aut nesciunt. Ducimus rem asperiores ut. Ducimus aliquam non ipsa dolorum non aut.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/2b5b7288-0999-4f21-a1be-4bcf756ac1d1/5858ebe8-cb8a-4939-abef-c9214b03caa4.png","created_on":1651783673,"modified_on":1651783673},{"id":"ef08c350-b538-4a86-ba50-cdda78e05ee4","title":"totam illo maiores","code":"totam-illo-maiores","description":"Incidunt id asperiores beatae sit similique deserunt quo. Ea accusantium perspiciatis atque sint quis nam reprehenderit. Error enim ut dolorem accusamus alias et aut. Reprehenderit et aliquam quisquam molestiae iusto.","created_on":1651783673,"modified_on":1651783673},{"id":"f3996f44-e226-4a7f-a4f3-70ea305cceae","title":"est nihil omnis","code":"est-nihil-omnis","description":"Explicabo voluptatem voluptatem voluptatem nesciunt nam sit. Est earum voluptas pariatur voluptatem. Quas et sint.","created_on":1651783673,"modified_on":1651783673},{"id":"9530bf39-ba45-4196-976b-50587e4e9da3","title":"sapiente reprehenderit quam","code":"sapiente-reprehenderit-quam","description":"Officia possimus officia consequuntur ipsum velit commodi quasi qui non. Sunt numquam rem non asperiores iste quasi quae eveniet. Iure est quis rerum pariatur doloremque in et est. Aut delectus inventore omnis aut deleniti omnis nihil. Sed mollitia qui molestiae occaecati quisquam. Qui est est.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/9530bf39-ba45-4196-976b-50587e4e9da3/75ea4c0f-3f05-40f9-bb77-b1c42a32f8bf.png","created_on":1651783673,"modified_on":1651783673},{"id":"41e60cfb-b103-4c67-ab15-768fb2137fd6","title":"eos pariatur possimus","code":"eos-pariatur-possimus","description":"Maxime sit quasi enim aut. Omnis nam veniam sint et delectus ipsa. Atque ea quia ipsum distinctio molestiae unde sed molestiae. Similique earum est. Optio eligendi quas hic iste quam aperiam. Reprehenderit alias commodi consequatur molestiae et sed.","created_on":1651783673,"modified_on":1651783673},{"id":"7a90f495-0acb-4b5d-a54d-951b55cc675d","title":"vel nihil debitis","code":"vel-nihil-debitis","description":"Eius non et fuga deserunt tempora placeat. Voluptatum ratione ipsa in voluptatem debitis atque laudantium pariatur. Omnis eaque quae ut vero eos non.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/7a90f495-0acb-4b5d-a54d-951b55cc675d/5560bd3a-3231-4392-a7aa-2c490dd937a6.png","created_on":1651783673,"modified_on":1651783673},{"id":"d5fce3f0-aed5-4f73-b820-18782f0d9d4d","title":"est rem ea","code":"est-rem-ea","description":"Earum autem sed atque rerum sed corrupti. A iure velit asperiores id quisquam. Id ratione veniam non laudantium dolore dolorem mollitia et ut. Ipsum omnis et nulla. Fuga odio commodi rerum consequuntur explicabo enim facere tempore. Non nesciunt consequatur aut.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/d5fce3f0-aed5-4f73-b820-18782f0d9d4d/4db1a487-cf97-4a1b-984e-49c514ace598.png","created_on":1651783673,"modified_on":1651783673},{"id":"cb870769-b078-49ff-9305-9b481ea040fc","title":"sed sit voluptas","code":"sed-sit-voluptas","description":"Distinctio sed aut nihil. Dolor earum dolorum. Ut commodi aut aliquam ea. Quam iusto sapiente. Deserunt et eaque et maiores facilis dolor nulla qui error. Dolorum necessitatibus provident dolores aliquid repellendus reprehenderit voluptas ratione et.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/cb870769-b078-49ff-9305-9b481ea040fc/88883ccd-4203-4e2d-8fe4-dcea56bc3df6.png","created_on":1651783673,"modified_on":1651783673},{"id":"c4970fa3-58b2-4114-9a61-a210b422f7f9","title":"quod nulla debitis","code":"quod-nulla-debitis","description":"Quia alias est. Corrupti laboriosam rem consectetur veritatis quis maxime cum. Sit quam maxime excepturi accusamus aut ratione sint iste nostrum. Et beatae molestiae.","created_on":1651783673,"modified_on":1651783673},{"id":"8769be58-6210-4b95-920a-496959896875","title":"eaque assumenda dolorem","code":"eaque-assumenda-dolorem","description":"Reiciendis culpa corrupti distinctio molestiae voluptas totam. Delectus eum commodi cupiditate corrupti reprehenderit iusto. Provident minus aut quae recusandae itaque tempora sunt maiores ex. Et consectetur alias sed accusamus vel voluptate.","created_on":1651783673,"modified_on":1651783673},{"id":"629b09e9-075e-44e0-9836-a80f1e7550ef","title":"quia quaerat porro","code":"quia-quaerat-porro","description":"Illo non illo dolorum nobis quibusdam nostrum libero rerum mollitia. Facere deserunt alias debitis blanditiis. Consequatur quam sint aut nostrum molestiae libero quod quia.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/629b09e9-075e-44e0-9836-a80f1e7550ef/aef4b97b-98be-4a51-a183-1025390ae91f.png","created_on":1651783673,"modified_on":1651783673},{"id":"aa579a69-edf4-4f71-9b5e-ec2f8489292b","title":"quia cumque non","code":"quia-cumque-non","description":"Libero dolores voluptas quidem exercitationem iure et debitis. Mollitia distinctio et. Repellat laudantium odio sed quo. A occaecati et qui quibusdam molestias architecto praesentium vero debitis. Rerum temporibus suscipit porro sequi. Impedit possimus illo sint est commodi ex consequuntur atque et.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/aa579a69-edf4-4f71-9b5e-ec2f8489292b/c11b3d47-8c90-48a6-8ac4-d29cf6f5319d.png","created_on":1651783673,"modified_on":1651783674}]}`
