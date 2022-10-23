<template>
  <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm"
    style="max-width:400px;">
    <el-form-item label="用户名" prop="username">
      <el-input v-model="ruleForm.username" type="text" autocomplete="off" />
    </el-form-item>

    <el-form-item label="密码" prop="pass">
      <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
    </el-form-item>


  </el-form>
</template>

<script lang="ts" setup>
import axios from 'axios'
import { reactive, ref } from 'vue'
import { FormInstance } from 'element-plus'
import { useRouter } from 'vue-router'
import { useMainStore } from '@/stores/UserState'

const userState = useMainStore()
const router = useRouter()


const ruleFormRef = ref<FormInstance>()

const ruleForm = reactive({
  username: '123',
  pass: '123',
  checkPass: '123',
  email: 'hello@qq.com'
})

const rules = reactive({
  username: [{ required: true, trigger: 'blur' }],
  pass: [{ required: true, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      let formData = {
        'username': ruleForm.username,
        'password': ruleForm.pass
      }
      axios.post(process.env.VUE_APP_BASE_API + '/user/login', formData).then(
        function (response) {
          if (response.data.code === 0) {
            alert("登录成功！跳转到个人主页")
            localStorage.setItem("userInfo", JSON.stringify(response.data.resp))
            userState.getUserState()
            router.push("all-books")
          } else {
            alert("登录失败！请检查输入")
          }
        }
      ).catch(
        function (err) {
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
