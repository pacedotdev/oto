//
//  main.swift
//  SwiftCLIExample
//
//  Created by Mat Ryer on 28/09/2020.
//

import Foundation

let client = OtoClient(withEndpoint: "http://localhost:8080/oto")
let greeterService = GreeterService(withClient: client)

greeterService.greet(withRequest: GreetRequest(
    name: "Mat"
)) { (response, err) -> () in
    if let err = err {
        print("ERROR: \(err)")
        return
    }
    print(response!.greeting!)
}

sleep(1)
