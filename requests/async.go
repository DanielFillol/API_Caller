package requests

import (
	"github.com/DanielFillol/API_Caller/models"
	"sync"
)

//AsyncAPIRequest uses go wait groups to run parallel routines on APIRequest, how many is set by numberOfWorkers.
//	all requests are organized by StreamInputs
//	all responses are returned as []models.WriteStruct to be uses by readFile function
func AsyncAPIRequest(users []string, email string, password string, numberOfWorkers int) ([]models.WriteStruct, error) {
	done := make(chan struct{})
	defer close(done)

	inputCh := StreamInputs(done, users)

	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	resultCh := make(chan models.WriteStruct)

	for i := 0; i < numberOfWorkers; i++ {
		// spawn N worker goroutines, each is consuming a shared input channel.
		go func() {
			for input := range inputCh {
				bodyStr, err := APIRequest(input, email, password)
				resultCh <- models.WriteStruct{
					SearchName:     bodyStr.SearchName,
					ID:             bodyStr.ID,
					CreatedAt:      bodyStr.CreatedAt,
					UpdatedAt:      bodyStr.UpdatedAt,
					DeletedAt:      bodyStr.DeletedAt,
					Name:           bodyStr.Name,
					Classification: bodyStr.Classification,
					Metaphone:      bodyStr.Metaphone,
					NameVariations: bodyStr.NameVariations,
					Err:            err,
				}
			}
			wg.Done()
		}()
	}

	// Wait all worker goroutines to finish. Happens if there's no error (no early return)
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []models.WriteStruct
	for result := range resultCh {
		if result.Err != nil {
			// return early. done channel is closed, thus input channel is also closed.
			// all worker goroutines stop working (because input channel is closed)
			return nil, result.Err
		}
		results = append(results, result)
	}

	return results, nil
}

//StreamInputs receives a slice of strings as inputs and a "done" channel to signal when to stop processing.
//	It returns a channel that streams each input string in the slice sequentially.
//	If the "done" channel is closed prematurely (due to an error midway), it closes the input channel and returns.
//	This function is designed to be used in a pipeline with other functions that process the input stream.
func StreamInputs(done <-chan struct{}, inputs []string) <-chan string {
	inputCh := make(chan string)
	go func() {
		defer close(inputCh)
		for _, input := range inputs {
			select {
			case inputCh <- input:
			case <-done:
				// in case done is closed prematurely (because error midway),
				// finish the loop (closing input channel)
				break
			}
		}
	}()
	return inputCh
}
