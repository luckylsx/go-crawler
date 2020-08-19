package persist

import (
	"context"
	"errors"
	"log"

	"imooc.com/learngo/crawler/engine"

	"github.com/olivere/elastic"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d: %v", itemCount, item)
			itemCount++
			err := save(index, client, item)
			if err != nil {
				log.Printf("Item Saver : error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(index string, client *elastic.Client, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("must apply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err
	}
	// fmt.Printf("%+v\n", resp)
	// output : &{Index:dating_profile Type:zhenai Id:rC7t8HMBQ2fIxEgKs7ro Version:1 Result:created Shards:0xc0002441c0 SeqNo:0 PrimaryTerm:1 Status:0 ForcedRefresh:false}
	return nil
}
