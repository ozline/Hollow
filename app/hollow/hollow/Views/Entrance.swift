//
//  Entrance.swift
//  hollow
//
//  Created by ozline on 2022/8/1.
//

import SwiftUI

struct EntranceView: View {
    
    @State var isLogin:Bool = false
    var body: some View {
//        return Group{
//            if isLogin { //登录成功进入主界面
//                ContentView()
//            }else{
//                LoginView(successLogin: $isLogin)
//            }
//        }
        return Group{
            MainView()
        }
    }
}
