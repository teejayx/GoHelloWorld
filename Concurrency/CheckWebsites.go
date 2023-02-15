package concurrency



type WebsiteChecker func(string) bool 

type result struct{
	string 
	bool 
}

//chanels are go data structure that can both recieve and send values, this operion along side their details, allow cmmmunicatin betwee  differenct processes 
// chan result is a type of channel - a channel of result 
// send statment resultChanel  <--- result{u, wc(u)}
//
//

func CheckWebsites(wc WebsiteChecker, urls []string) map[string] bool{
	results := make(map[string]bool)
	resultChannel := make(chan result)

     for _, url := range urls {
		go func (u string)  {
			resultChannel <- result{u, wc(u)}
		}(url)
			
	 }

	 for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	 }

	 return results
}