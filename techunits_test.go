package main

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-examples-client/techunits"
)

/*

type mockDepositServer struct {
	pb.UnimplementedDepositServiceServer
}

func (*mockDepositServer) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	if req.GetAmount() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot deposit %v", req.GetAmount())
	}

	return &pb.DepositResponse{Ok: true}, nil
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterDepositServiceServer(server, &mockDepositServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

*/

// var entryLists [25]pb.EntryResponse

// func TestMain(m *testing.M) {
// 	json.Unmarshal([]byte(entryList), &entryLists)
// }

func TestEntryListRequest_ListEntries(t *testing.T) {
	tests := []struct {
		name string
		res  string
		err  error
	}{
		{
			"default",
			entryList,
			nil,
		},
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleServiceClient(conn)

	// Contact the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

const entryList = `{"entries":[{"id":"5eee2c18-74e2-414f-b2cd-6e0e3e691aba","title":"voluptate alias libero","code":"voluptate-alias-libero","description":"Ipsam et incidunt sit tenetur sit ut et quis quis. Autem ad dolores nesciunt impedit earum et quod. Commodi ut enim molestiae voluptas.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/5eee2c18-74e2-414f-b2cd-6e0e3e691aba/b7252a0a-6214-4c9f-aa26-bb8cdb18287b.png","created_on":1651770172,"modified_on":1651770172},{"id":"b743c62a-56fb-4896-9742-2d013512a552","title":"labore et voluptas","code":"labore-et-voluptas","description":"Reprehenderit aut vitae ipsum voluptas veniam omnis unde distinctio ut. Ex molestias consequatur sequi molestiae consectetur totam et. Porro qui delectus quidem veniam perspiciatis illo vero. Porro exercitationem est iusto. Labore recusandae distinctio quibusdam debitis et autem neque molestias molestiae. Eius dolorum ut iste qui ut voluptates vel autem.","created_on":1651770172,"modified_on":1651770172},{"id":"6d4d832d-8541-4e2b-bca3-7e95d1aeb0db","title":"animi doloribus aliquam","code":"animi-doloribus-aliquam","description":"Sunt temporibus laboriosam sint excepturi nemo facilis quos est. Amet qui aut voluptas qui corporis. Possimus quae et et dolor suscipit veniam velit. Porro eos enim et exercitationem sapiente et nobis rerum nostrum. Qui corporis repellendus placeat asperiores ullam esse facere reprehenderit. Iure quos quibusdam et explicabo in voluptate laboriosam quidem repellat.","created_on":1651770172,"modified_on":1651770172},{"id":"5b9c9600-f81d-4494-8043-e4ae21b53860","title":"dolores voluptates facere","code":"dolores-voluptates-facere","description":"Dolor laudantium esse. Vero quia et. Neque debitis quasi. Et molestias sit aliquam totam et aut velit exercitationem.","created_on":1651770172,"modified_on":1651770172},{"id":"109dc19a-df0f-4263-bf6e-bd6d05f688ba","title":"ut nisi quae","code":"ut-nisi-quae","description":"Beatae tenetur neque possimus accusantium est sed laudantium rem. Vero molestiae ut. Laborum inventore modi. Vel ratione laudantium sunt fugit odio et corrupti et.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/109dc19a-df0f-4263-bf6e-bd6d05f688ba/6859053b-a1d1-411c-9c09-4c0c1498be3b.png","created_on":1651770172,"modified_on":1651770172},{"id":"d6476e30-b1a8-4d5a-919d-2df9a2822c67","title":"sapiente quia impedit","code":"sapiente-quia-impedit","description":"Cum id omnis illum. Ab aspernatur totam dolores officia. Adipisci magnam at rerum accusamus qui fugiat. Dolores doloribus molestiae. Magni sit dolorem ex omnis. Praesentium et voluptatem soluta vel.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/d6476e30-b1a8-4d5a-919d-2df9a2822c67/62f66ad8-424d-4209-b913-df88ec4aec6b.png","created_on":1651770172,"modified_on":1651770172},{"id":"6002844e-2897-4b9c-b03d-4d0f4d01f84d","title":"eius facere asperiores","code":"eius-facere-asperiores","description":"Autem in quam et aspernatur voluptatem architecto iusto velit mollitia. Deserunt vero sunt quod ipsam. Iusto sint non in earum laudantium eligendi.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/6002844e-2897-4b9c-b03d-4d0f4d01f84d/bc0198f3-d8f5-48c4-bd5c-ea965f8a18e4.png","created_on":1651770172,"modified_on":1651770172},{"id":"6febd181-120f-4c07-9e8f-07a25acbad40","title":"quos expedita temporibus","code":"quos-expedita-temporibus","description":"Magni quasi asperiores architecto minus laudantium voluptas. Repellendus dolores soluta molestiae facilis quos. Hic est beatae ut eum id beatae eum ducimus. Iure dolor blanditiis molestias non iure aspernatur. Modi fugit et facilis sed necessitatibus minus.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/6febd181-120f-4c07-9e8f-07a25acbad40/6e2073f6-24ac-44a8-b3ac-7c57d69efb0c.png","created_on":1651770172,"modified_on":1651770172},{"id":"3186e7f1-cb1d-4d91-8ab6-a441e831b870","title":"exercitationem voluptatem tempore","code":"exercitationem-voluptatem-tempore","description":"Et quia et quia ab ut et. Impedit earum consequuntur ut illo consequatur. Ducimus magnam quia ipsam.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/3186e7f1-cb1d-4d91-8ab6-a441e831b870/3946bc0b-7ce0-4677-8172-0bbe77f44e79.png","created_on":1651770172,"modified_on":1651770172},{"id":"d55771f7-1205-4b0f-b443-edf5c22a0aff","title":"est et assumenda","code":"est-et-assumenda","description":"Qui ratione ipsa voluptatum non quis quia laborum veniam. Modi quis sed. Vero ipsa perspiciatis in. Sit et ut et similique dolorum ullam rerum. Omnis enim qui.","created_on":1651770172,"modified_on":1651770172},{"id":"e20ead34-1573-4580-b3be-c8ef9daec238","title":"assumenda et cum","code":"assumenda-et-cum","description":"Itaque exercitationem molestiae et omnis consectetur aliquid pariatur aut excepturi. Et consequatur rerum id dolores omnis. Harum sint molestiae. Qui qui officiis enim magni consequuntur ab maxime reiciendis est. Porro aut quasi sapiente earum natus totam deleniti.","created_on":1651770172,"modified_on":1651770172},{"id":"bc1478a4-4890-41e1-a84b-edfe5d8274cc","title":"qui sit necessitatibus","code":"qui-sit-necessitatibus","description":"Voluptas hic placeat est in provident. Natus praesentium aut consequatur culpa. Sunt nihil reiciendis. Sed accusantium accusantium omnis sequi facere beatae rerum. Incidunt fugit optio repellat illo. Explicabo inventore non fugit nam asperiores facilis placeat esse ut.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/bc1478a4-4890-41e1-a84b-edfe5d8274cc/95afbb64-6473-430c-ad3f-5d7c4c405faf.png","created_on":1651770172,"modified_on":1651770172},{"id":"060b528c-593a-4260-ac37-9150147f1c89","title":"sed nulla ipsam","code":"sed-nulla-ipsam","description":"Consequuntur qui a. Ut nam debitis voluptatem impedit magni. Quasi sed repellendus consequatur. Laborum pariatur quod ut. Facilis enim voluptatibus ex.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/060b528c-593a-4260-ac37-9150147f1c89/9bf3adf9-e96e-49d6-9fc4-a03537f01495.png","created_on":1651770172,"modified_on":1651770172},{"id":"383eca7d-a488-4296-895f-d5ba3bd02922","title":"quibusdam qui provident","code":"quibusdam-qui-provident","description":"Aut et et quaerat eum eaque temporibus. Architecto sint autem beatae ad quo vitae esse. Eveniet minima in perferendis. Molestiae sed ut similique animi et ratione. Ducimus voluptatem omnis odio. Quisquam explicabo amet quis tempora reiciendis magni quae quam ut.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/383eca7d-a488-4296-895f-d5ba3bd02922/b1258add-f9ce-4fed-a6f3-fb840ba39c69.png","created_on":1651770172,"modified_on":1651770172},{"id":"715aed44-85d3-47df-840a-a5de271c913a","title":"ut harum quia","code":"ut-harum-quia","description":"Quia dicta rerum a. Laudantium cumque magni. Pariatur reiciendis doloremque facilis quod laboriosam. Molestiae ipsam voluptatum ratione. Aut non expedita enim et fuga et dolor.","created_on":1651770172,"modified_on":1651770172},{"id":"9d11c690-3e1f-40ba-b593-343d4b870e52","title":"eum quam aut","code":"eum-quam-aut","description":"Temporibus eum nesciunt accusantium vitae consequuntur eligendi. Non ea sint fuga perferendis aspernatur necessitatibus. Placeat at cum eligendi esse beatae quod. Rem et sunt voluptas numquam sit et dolore. Sequi aliquam consequatur reiciendis sunt dolorem debitis modi est reiciendis. Provident eius enim voluptatem provident molestiae iure nostrum labore.","created_on":1651770172,"modified_on":1651770172},{"id":"ffff0bc9-233e-46cf-8975-5ab9c5fea6a9","title":"voluptatem similique vero","code":"voluptatem-similique-vero","description":"Saepe rerum et doloribus exercitationem magnam vitae debitis id. Dolorem animi libero nam minima ipsam suscipit quo. Quos non quasi quisquam laudantium libero eos adipisci necessitatibus. Ea aut neque enim quibusdam temporibus eum pariatur sit numquam.","created_on":1651770172,"modified_on":1651770172},{"id":"6398b32a-5fec-4fe5-bc18-df0863e6c8bc","title":"commodi consequatur et","code":"commodi-consequatur-et","description":"Totam aut aliquam distinctio tempora exercitationem animi explicabo. Iusto quia magni voluptatem quia exercitationem. Ducimus omnis ducimus ea. Eligendi et ipsam et laudantium porro et voluptates nemo iusto. Iure et corporis sapiente nostrum est unde excepturi.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/6398b32a-5fec-4fe5-bc18-df0863e6c8bc/d3a3a045-b21d-43e6-b013-bd8571f0da31.png","created_on":1651770172,"modified_on":1651770172},{"id":"29bd32c1-e82e-4bb5-aeee-24e005c03a8f","title":"optio vel atque","code":"optio-vel-atque","description":"Magnam omnis repellat et iusto. Blanditiis debitis velit. Sapiente eos voluptate eos rerum et ratione. Voluptas velit consequuntur quos. Voluptatem voluptate dignissimos aut nemo nobis occaecati et suscipit nam.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/29bd32c1-e82e-4bb5-aeee-24e005c03a8f/49b6e4f3-2ff4-4aee-b54e-ce41bb691f54.png","created_on":1651770172,"modified_on":1651770172},{"id":"47a0174f-b62b-4b5b-aa61-c9e6675fe2a3","title":"dolorem nemo quia","code":"dolorem-nemo-quia","description":"Atque magnam labore quos omnis in quas. Provident natus beatae enim. Autem optio impedit. Non vel eaque et eum ratione voluptas eum.","created_on":1651770172,"modified_on":1651770172},{"id":"1b880d38-c07b-426d-b77c-eb25fe6f8178","title":"porro odio dolore","code":"porro-odio-dolore","description":"Repellendus ut laborum rerum pariatur incidunt eos. Deleniti molestiae eligendi. Et quia eligendi a sit et nisi. At rem et magni autem omnis vero illo minus. Aliquid totam nihil sint. Et nihil necessitatibus sint ad.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/1b880d38-c07b-426d-b77c-eb25fe6f8178/118a0745-3a53-4503-b115-6ba708060b0d.png","created_on":1651770172,"modified_on":1651770172},{"id":"4323430f-d6ff-4416-9991-57d5d85fd5a0","title":"molestiae libero officia","code":"molestiae-libero-officia","description":"Amet iusto quia qui corrupti praesentium possimus accusantium necessitatibus. Et deserunt est. Accusamus accusantium voluptatem qui alias error et quia voluptates.","created_on":1651770172,"modified_on":1651770172},{"id":"9b1f1036-0748-4659-a41a-0ae7d7fc33b7","title":"iusto numquam blanditiis","code":"iusto-numquam-blanditiis","description":"Corrupti at sint. Inventore nostrum illo. Sequi sit et nesciunt accusantium atque odit dignissimos vel libero.","created_on":1651770172,"modified_on":1651770172},{"id":"bf8aeeaf-6b08-4bfb-a217-ad61b84b1e4f","title":"officiis facere adipisci","code":"officiis-facere-adipisci","description":"Est dolores magnam illum debitis. Praesentium quam occaecati facere doloremque et accusamus. Odit voluptas officiis neque eos quam earum. Minus voluptas voluptas nisi delectus officiis. Vel officia ea quasi et.","created_on":1651770172,"modified_on":1651770172},{"id":"8be7e34e-9f32-4922-b0fc-ba7698a3fe43","title":"cumque ea recusandae","code":"cumque-ea-recusandae","description":"Sunt nemo voluptatum impedit eum qui tenetur. Quo adipisci eaque nemo eum consequatur numquam laboriosam quia est. Sunt dolore dolorem voluptatibus ut. Nulla voluptas maxime corporis non incidunt voluptate quo. Qui molestiae fugiat voluptatum et voluptatem.","file_url":"https://grpc-techunits.ams3.digitaloceanspaces.com/uploads/8be7e34e-9f32-4922-b0fc-ba7698a3fe43/a61375ec-7a87-483c-9464-380765ce6670.png","created_on":1651770172,"modified_on":1651770172}]}`
