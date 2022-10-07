<template>
  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    status-icon
    :rules="rules"
    label-width="120px"
    class="demo-ruleForm"
    style="max-width:400px;"
  >
    <el-form-item label="用户名" prop="username">
      <el-input v-model="ruleForm.username" type="text" autocomplete="off" />
    </el-form-item>

    <el-form-item label="密码" prop="pass">
      <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
    </el-form-item>

    <el-form-item label="确认密码" prop="checkPass">
      <el-input
        v-model="ruleForm.checkPass"
        type="password"
        autocomplete="off"
      />
    </el-form-item>

    <el-form-item label="邮箱" prop="email">
      <el-input v-model="ruleForm.email" />
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="submitForm(ruleFormRef)">注册</el-button>
    </el-form-item>

  </el-form>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { FormInstance } from 'element-plus'
import axios from 'axios'


const ruleFormRef = ref<FormInstance>()


const validatePass = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input the password'))
  } else {
    if (ruleForm.checkPass !== '') {
      if (!ruleFormRef.value) return
      ruleFormRef.value.validateField('checkPass', () => null)
    }
    callback()
  }
}
const validatePass2 = (rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('Please input the password again'))
  } else if (value !== ruleForm.pass) {
    callback(new Error("Two inputs don't match!"))
  } else {
    callback()
  }
}

const ruleForm = reactive({
  username: '123',
  pass: '123',
  checkPass: '123',
  email: 'hello@qq.com'
})

const rules = reactive({
  username: [{required: true, trigger: 'blur'}],
  pass: [{ validator: validatePass, trigger: 'blur' }],
  checkPass: [{ validator: validatePass2, trigger: 'blur' }],
  email: [{type: 'email', required: true, trigger: 'blur'}]
})

const submitForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      console.log('submit!')
      let formData = new FormData()
      formData.append('username', ruleForm.username)
      formData.append('password', ruleForm.pass)
      formData.append('email', ruleForm.email)
      const url = 'http://localhost:8081/register'
      axios.post(url, formData).then(
        function(response){
          console.log(response.data)
        }
      ).catch(
        function(err) {
          console.log(err.message)
        }
      )
    } else {
      console.log('error submit!')
      return false
    }
  })
}
</script>
