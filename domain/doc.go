// Package domain contains the base types and interfaces used to develop on the Pii platform
//
//
//
//
// e.g. This is a simple processor that pulls new line delimited strings from a URL, loads those strings into a package
// variable, and when called to Process, takes the data and checks if it is in the list of strings
//
//   package main
//
//   import (
//       "io/ioutil"
//       "net/http"
//       "strings"
//
//       "github.com/tyndyll/lintol-server/domain"
//   )
//
//   func init() {
//       // disregard the errors for demo purposes
//       resp, _ := http.Get("http://example.org/pii-terms")
//       data, _ := ioutil.ReadAll(resp.Body)
//       listOfTerms = strings.Split("\n")
//   }
//
//   var listOfTerms = []string{}
//
//   var helpInfo = *domain.HelpInformation{
//       Name: "demo-processor",
//       Description: "Checks if your data contains one of the prescribed strings"
//       Options: map[string]string
//   }
//
//   type DemoProcessor struct {}
//
//   func (processor *DemoProcessor) Process(data *string, config *map[string]interface{}) (*domain.ProcessorResult, error) {
//       results = []string{}
//       for _, term := range listOfTerms {
//           if strings.Contains(data, term) {
//               results = append(results, term)
//           }
//       }
//       if len(results) == 0 {
//           return nil, nil
//       }
//       return &domain.ProcessorResult{
//           Name: helpInfo.Name,
//           Results: results
//       }, nil
//   }
//
//   func (processor *DemoProcessor) Help() *domain.HelpInformation {
//       return helpInfo
//   }
//
//   // Explicitly output the Processor
//   var Processor *processors.IPAddress
package domain

