//
//  MainView.swift
//  hollow
//
//  Created by ozline on 2022/8/28.
//

import SwiftUI
import CoreData

struct MainView: View {
    let persistenceController = PersistenceController.shared
    
    @State private var username:String = ""
    @State private var password:String = ""
    @State private var loginActive:Bool = false
    @State private var selection:Int = 0
    
    var handler: Binding<Int>{
        Binding(
            get: {
                self.selection
            },
            set: {
                self.selection = $0
            }
        )
    }
    
    init(){
        var tmp : [String : String] = ["page" : "1","pagesize" : "10"]
        unifiedRequest(url: "http://localhost:9001/api/forest/all", method: .get, params: tmp)
//        UITabBar.appearance().backgroundColor = .lightGray
    }
    
    var body: some View {
        TabView(selection: handler){
            //主页
//            ItemListView()
//            .tabItem({
//                Image(systemName: "house")
//                Text("首页")
//            })
//            .tag(0)
//
//            //搜索
//            ItemSearchView()
//                .tabItem({
//                    Image(systemName: "equal")
//                    Text("搜索")
//                })
//                .tag(3)
//
//            //个人资料页
//            ProfileView(userid:Global.userid)
//            .tabItem({
//                    Image(systemName: "person")
//                    Text("我的")
//                })
//            .tag(2)
        }
        .onAppear(){
        }
    }
}
