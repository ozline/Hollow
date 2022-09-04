//
//  hollowApp.swift
//  hollow
//
//  Created by ozline on 2022/8/1.
//

import SwiftUI

@main
struct lazyfishApp: App {
    
    var tmp : [String : String] = ["page" : "1","pagesize" : "10"]

    var body: some Scene {
        WindowGroup {
            EntranceView()
        }
    }
}
