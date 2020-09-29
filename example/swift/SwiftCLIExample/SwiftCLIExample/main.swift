//
//  main.swift
//  SwiftCLIExample
//
//  Created by Mat Ryer on 28/09/2020.
//

import Foundation

let client = OtoClient(withEndpoint: "http://localhost:8080/oto")
let service = MyService(withClient: client)

service.doSomething(withRequest: DoSomethingRequest(
    name: "Mat"
)) { (response, err) -> () in
    print("done")
}

class OtoClient {
    var endpoint: String
    init(withEndpoint url: String) {
        self.endpoint = url
    }
}

class MyService {
    var client: OtoClient
    init(withClient client: OtoClient) {
        self.client = client
    }
    func doSomething(withRequest doSomethingRequest: DoSomethingRequest, completion: (_ response: DoSomethingResponse?, _ error: Error?) -> ()) {
        //var request = URLRequest(url: URL(string: "\(self.client.endpoint)/MyService/MyMethod")!)
        // https://jsonplaceholder.typicode.com/todos/1
        var request = URLRequest(url: URL(string: "https://jsonplaceholder.typicode.com/todos/1")!)
        
        request.httpMethod = "POST"
        request.addValue("application/json; charset=utf-8", forHTTPHeaderField: "Content-Type")
        request.addValue("application/json; charset=utf-8", forHTTPHeaderField: "Accept")
        var jsonData: Data
        do {
            jsonData = try JSONEncoder().encode(doSomethingRequest)
        } catch let jsonEncodeErr {
            print("TODO: handle JSON encode error: \(jsonEncodeErr)")
            return
        }
        request.httpBody = jsonData
        let session = URLSession(configuration: URLSessionConfiguration.default)
        let task = session.dataTask(with: request) { (data, response, error) in
            if let err = error {
                print("TODO: handle response error: \(err)")
                return
            }
            var doSomethingResponse: DoSomethingResponse
            do {
                doSomethingResponse = try JSONDecoder().decode(DoSomethingResponse.self, from: data!)
            } catch let err {
                print("TODO: handle JSON decode error: \(err)")
                return
            }
            print("\(doSomethingResponse)")
        }
        task.resume()
    }
}

struct DoSomethingRequest: Encodable {
    var name: String = ""
}

struct DoSomethingResponse: Decodable {
    var greeting: String = ""
}
