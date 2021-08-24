import COS from 'cos-js-sdk-v5'
import { getTmpSecret } from '@/api/z_user'
import { Message } from 'element-ui'

export default function upload(file,callback) {
    getTmpSecret().then(response => { // 后台接口返回 密钥相关信息
        const data = response.data
        let uploadFileName = Date.parse(new Date())
        let cos = new COS({
            getAuthorization: function(options, callback) {
                callback({
                    TmpSecretId: data.tmp_secret_id,
                    TmpSecretKey: data.tmp_secret_key,
                    XCosSecurityToken: data.session_token,
                    StartTime: data.start_time,
                    ExpiredTime: data.expired_time,
                    expiration: data.expiration,
                    requestId: data.request_id
                })
            }
        })
        cos.putObject(
            {
                Bucket: data.bucket_name, // 存储桶名称
                Region: 'ap-chengdu', // 地区
                Key: '/tweet/img/' + uploadFileName + '.png', // 图片名称
                StorageClass: 'STANDARD',
                Body: file, // 上传文件对象
                onHashProgress: function (progressData) {
                    console.log ('校验中', JSON.stringify (progressData));
                },
                onProgress: function (progressData) {
                    console.log ('上传中', JSON.stringify (progressData));
                },
            },
            function (err, data) {
                console.log('999',err,data,data.Location)
                if (err) {
                    Message({message:'文件上传失败,请重新上传',type: 'error',})
                } else {
                    let fileUrl = 'http://' + data.Location
                    callback(fileUrl)
                }
            }
        )
    })
}