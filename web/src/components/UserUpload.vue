/* eslint-disable @typescript-eslint/no-non-null-assertion */
<template>
<div style="max-width:400px;">
  <el-form
    ref="booInfoFormRef"
    :model="bookInfo"
    status-icon
    label-width="120px"
    class="demo-ruleForm"
    style="max-width:400px;"
  >
    <el-form-item label="上传文件" prop="file">
      <!-- name是上传的文件字段名，也就是后端接收的form索引的key，默认是file -->
      <el-upload
        ref="upload"
        class="upload-demo"
        :action="api_url"
        :limit="1"
        :on-exceed="handleExceed"
        :auto-upload="false"
        :on-success="handleSuccess"
        name="file"
      >
        <template #trigger>
          <el-button type="primary">select file</el-button>
        </template>
        <el-button class="ml-3" type="success" @click="submitUpload">
          upload to server
        </el-button>
        <template #tip>
          <div class="el-upload__tip text-red">
            limit 1 file, new file will cover the old file
          </div>
        </template>
      </el-upload>
    </el-form-item>
    <el-form-item label="书名" prop="bookName">
      <el-input v-model="bookInfo.bookName" type="text" />
    </el-form-item>

    <el-form-item label="价格" prop="bookPrice">
      <el-input 
      v-model="bookInfo.bookPrice" 
      type="text" 
      oninput="value=value.replace(/[^0-9.]/g,'')"
      />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm(booInfoFormRef)">新增</el-button>
    </el-form-item>
  </el-form>
</div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { genFileId } from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile, FormInstance } from 'element-plus'
import axios from 'axios'

const upload = ref<UploadInstance>()
const api_url = ref(process.env.VUE_APP_BASE_API+"/book/upload")
const booInfoFormRef = ref<FormInstance>()
const bookInfo = reactive({
  userID: '',
  bookName: '',
  bookPrice: '',
  picURL: ''
})
const usesrInfo = JSON.parse(localStorage.getItem("userInfo")!)    // 非空断言，叹号

const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()    // 这里的非空断言在项目根目录下的.eslintrc.js的rules中关掉，"@typescript-eslint/no-non-null-assertion": "off"
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)    // ! 断言了其前面的变量不为null，但是不知道为什么这里会有波浪线。
                                     // 把 ! 非空断言改为if判断就不再波浪线提示了
                                     // 上面写了原因，eslint的检测规则问题，关闭这一项检测就行了
                                     // 官方建议是不要用非空断言，所以这里给了波浪线
}

const submitUpload = () => {
  upload.value!.submit()
  
}

// 文件上传成功时的钩子
const handleSuccess = (res: any) => {
  // res即文件上传成功后的服务器返回值
  bookInfo.picURL = res.resp.PicURL
}

// 提交输入的信息到服务器
const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    // 检查是否已经完成图片的上传
    if (valid && bookInfo.picURL !== "") {
      let formData = {
        user_id: usesrInfo.username,
        name: bookInfo.bookName,
        price: parseFloat(bookInfo.bookPrice),
        pic_url: bookInfo.picURL
      }
      axios.post(process.env.VUE_APP_BASE_API+'/book/add', formData, {params:{token: usesrInfo.token}}).then(
        function(response) {
          if (response.data.code === 0) {
            alert("添加成功！")
          } else {
            alert("添加失败！请检查输入")
          }
        }
      ).catch(
        function(err) {
          alert("未知错误: " + err.message)
        }
      )
    } else {
      alert("请检查输入！")
      return false
    }
  })
}
</script>