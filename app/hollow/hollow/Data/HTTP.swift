//
//  HTTP.swift
//  hollow
//
//  Created by ozline on 2022/8/28.
//

import Foundation
import HandyJSON
import Alamofire

enum ReponseError : Int {
    case unknow = 0
    case success = 200
    case failure = 500
    case expired = 401
    case beOffLine = 402
}

struct ResponseModel<T:HandyJSON>{
    var errorCode       : ReponseError = .unknow
    var errorMessage    : String = "未知错误"
    var model           : T?
    var models          : [T?]?
    var resultData      : Any?
}

struct RequestHeaders : HandyJSON{
    var timestamp   : String?
    var token       : String?
    var sign        : String?
}

struct ResponseDefault : HandyJSON {}

struct ResponseData : HandyJSON{
    var code    : Int?
    var msg     : String?
    var data    : Any?
}

class NetManager {
    
    public class var `default` : NetManager {
        struct Static {
            static let instance : NetManager = NetManager()
        }
        
        return Static.instance
    }
}

//var Headers : HTTPHeaders{
//    get{
//        var headers = RequestHeaders()
//        headers.sign = ""
//        headers.token = ""
//        headers.timestamp = ""
//
//        guard let jsonHeader = headers.toJSON() , let jsonHeaders = jsonHeader as? [String:String] else {
//            return []
//        }
//
//        return HTTPHeaders.init(jsonHeaders)
//    }
//
//    // 接口地址
//    let RequestUrlHost : String = "http://localhost:8080/"
//
//    let ParameterEncoder : ParameterEncoder = URLEncodedFormParameterEncoder.default
//}

let prefix = "/apis/"

typealias RequestSuccess = (_  code : Int, _ msg : String, _ data : AnyObject) -> Void
//typealias RequestFail = (_ error : AnyObject) -> Void
//
//class NetWorkingService {
//
//    public func unifiedRequest(url : String,
//                               method: HTTPMethod,
//                               params: [String : Any]?,
//                               success: @escaping RequestSuccess,
//                               fail: @escaping RequestFail) {
//
//        var urlPath : String = url
//
//        if urlPath.contains(" "){
//            urlPath.replaceOccurrences(of: " ", with: "")
//        }
//
//        var tempParams = params
//        if tempParams == nil {
//            tempParams = [String : Any] ()
//        }
//
//        AF.request(_ convertible: URLConvertible,
//                   method: HTTPMethod = .get,
//                   parameters: Parameters? = nil,
//                   encoder: ParameterEncoder = URLEncodedFormParameterEncoder.default,
//                   headers: HTTPHeaders? = nil,
//                   interceptor: RequestInterceptor? = nil){
//        }
//    }
//
//}

func unifiedRequest<T: Encodable>(url: String,
                    method: HTTPMethod,
                    params: [String : T]?
                    ){
    AF.request(url,method: method,parameters: params,encoder: JSONParameterEncoder.default).response {
        response in
        
        debugPrint(response)
    }
}
