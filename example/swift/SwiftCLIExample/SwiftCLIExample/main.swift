//
//  main.swift
//  SwiftCLIExample
//
//  Created by Mat Ryer on 28/09/2020.
//

import Foundation

let client = OtoClient(withEndpoint: "http://localhost:8080/oto")
let service = MyService(withClient: client)

let resp = service.doSomething(request: DoSomethingRequest(
    name: "Mat"
))
print(resp.greeting)

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
    func doSomething(request: DoSomethingRequest) -> DoSomethingResponse {
        let resp = DoSomethingResponse(
            greeting: "Hi \(request.name)"
        )
        return resp
    }
}

struct DoSomethingRequest {
    var name: String = ""
}

struct DoSomethingResponse {
    var greeting: String = ""
}
