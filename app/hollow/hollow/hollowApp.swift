//
//  hollowApp.swift
//  hollow
//
//  Created by ozline on 2022/8/1.
//

import SwiftUI

@main
struct hollowApp: App {
    let persistenceController = PersistenceController.shared

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(\.managedObjectContext, persistenceController.container.viewContext)
        }
    }
}
