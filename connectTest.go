package main

import (
    "fmt"
    // "os"
    "encoding/json"

    "cgt.name/pkg/go-mwclient"
    "cgt.name/pkg/go-mwclient/params"
)

func main() {
    // Initialize a *Client with New(), specifying the wiki's API URL
    // and your HTTP User-Agent. Try to use a meaningful User-Agent.
    w, err := mwclient.New("http://bots.snpedia.com/api.php", "myWikibot")
    if err != nil {
        panic(err)
    }

    parameters := params.Values{
        "list": "categorymembers",
        "cmtitle": "Category:Is_a_snp",
    }

    // Print out the http request and response to screen
    // w.SetDebug(os.Stderr)
    q := w.NewQuery(parameters) // w being an instantiated Client
    queryRaw,err := w.GetRaw(parameters) // w being an instantiated Client

    // Decode raw json 
    // fmt.Println("queryRaw:",queryRaw)
    var dat map[string]interface{}
    json.Unmarshal(queryRaw, &dat)
    // fmt.Println("decodedQuery:",dat)
    // fmt.Println("\n\n")

    // Pass in a slice of SNP names
    // snpNames := []string{"Rs53576","Rs1815739"}
    // w.GetPagesByName(snpNames...)

    // w.GetPagesByID("4105","9994")

    // Define dat for decoding json into GO
    // var dat map[string]interface{}


    for i:=0; i < 5; i++{
        q.Next();
        // return the json response
        /*
        the query key has categorymemebers fields which itself is an array of objects with the following fields:
        we get the pageid and title for each snp

        "query": {
            "categorymembers": [
            {
                "ns": 0,
                "pageid": 10244,
                "title": "I1000001"
            },
      */
        response := q.Resp();
        // fmt.Println(response)
        fmt.Println("\n\n")

        // return the value of response in GO format which in this case is a map and categorymembers is an array of maps for the first page returned (jason package)
        /*
        &{map[categorymembers:
            [map[ns:0 pageid:10244 title:I1000001] map[ns:0 pageid:13450 title:I1000003] map[ns:0 pageid:19115 title:I1000004] map[ns:0 pageid:12979 title:I1000015] map[ns:0 pageid:13973 title:I3000001] map[ns:0 pageid:19671 title:I3000007] map[ns:0 pageid:19201 title:I3000014] map[ns:0 pageid:19667 title:I3000021] map[ns:0 pageid:19195 title:I3000029] map[ns:0 pageid:19177 title:I3000033]]] true}
        */
        // var members, err = response.GetObject("query");
        var members, _ = response.GetObjectArray("query","categorymembers");

        // Loop over members array
        for _, member := range members {
            var title, _ = member.GetString("title")
            fmt.Println(title)
        }       

    }

}