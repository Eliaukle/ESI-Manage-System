<template>
    <div class="login-pane">
      <n-card hoverable>
          <div style="text-align:center;">
            <n-icon size="20" color="#0e7a0d">
              <BarChart></BarChart>
            </n-icon>
            <span style="font-weight: bold; font-size: 20px;">&#8197;ESI排名管理系统</span>
          </div>
          <n-divider />
          <n-form :rules="rules" :model="user">
            <n-form-item-row path="userName" label="用户名">
              <n-input v-model:value="user.userName" placeholder="请输入用户名" />
            </n-form-item-row>
            <n-form-item-row path="password" label="密码">
              <n-input v-model:value="user.password" type = "password" placeholder="请输入密码" />
            </n-form-item-row>
          </n-form>
            <n-checkbox v-model:checked="user.rember" label="记住密码" />
          <n-button @click="login" type="primary" style="margin-top: 20px;" block secondary strong>
            登录
          </n-button>
      </n-card>
    </div>
  </template>
  
  <script setup>
  import {ref,reactive,inject} from 'vue'
  import {UserStore} from '../stores/UserStore'
  import {BarChart} from "@vicons/ionicons5"
  
  import {useRouter, useRoute} from 'vue-router'
  const router = useRouter()
  const route = useRoute()
  
  const axios = inject("axios")
  const message = inject("message")
  const userStore = UserStore()
  
  let rules = {
    userName: [
      { required: true, message: "请输入用户名", trigger: "blur" },
      { max: 20, message: "用户名长度小于 20 个字符", trigger: "blur"},
    ],
    password: [
      { required: true, message: "请输入密码", trigger: "blur" },
      { max: 20, message: "密码长度在小于 20 个字符", trigger: "blur"},    
    ]
  };
  
  const user = reactive({
    userName:localStorage.getItem("userName") || "",
    password:localStorage.getItem("password") || "",
    rember:localStorage.getItem("rember") == 1 || false
  })
  
  const login = async() => {
    let res = await axios.post("/login", {
      userName: user.userName,
      password: user.password
    })
    console.log(res)
    if (res.data.code == 200) {
        userStore.token = res.data.data.token
        if (user.rember) {
            localStorage.setItem("userName", user.userName)
            localStorage.setItem("password", user.password)
            localStorage.setItem("rember", user.rember? 1: 0)
        } else {
            localStorage.removeItem("userName")
            localStorage.removeItem("password")
            localStorage.setItem("rember", user.rember? 1: 0)
        }        
        router.push("/dashboard")
        message.success(res.data.msg)        
    } else {
        message.error(res.data.msg)
    }
  } 
  
  </script>
  
  <style lang="scss" scoped>
  .login-pane{
    width: 500px;
    margin: 0 auto;
    margin-top: 130px;
  }
  
  </style>