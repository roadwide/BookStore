import { reactive } from 'vue'
import axios from 'axios'

export function useGetBookList(UserID ?: string) {
    const data = {user_id:''}
    if (UserID !== undefined) {
        data['user_id'] = UserID
    }
    const bookInfo = reactive<any>([])
    axios.post(process.env.VUE_APP_BASE_API+'/book/info', data).then(
        function(response) {
            if (response.data.code === 0) {
                for(const oneBookData of response.data.resp) {    // of是内容，in是索引
                    bookInfo.push(oneBookData)
                }
            } else {
                alert("获取所有书目信息失败！")
            }
    }).catch(
        function(err) {
            alert("未知错误: " + err.message)
    })
    return bookInfo
}