<template>
    <div>
        <div class="container">
            <div class="topbar">
                <div style="line-height:50px;">
                    &#12288;&#12288;
                    <n-icon size="20" color="#0e7a0d">
                        <BarChart></BarChart>
                    </n-icon>
                    <span style="font-weight: bold; font-size: 20px;">&#8197;ESI排名管理系统</span>
                    <div style="position: absolute; top: 0px; right: 180px; font-size: 16px;">
                        欢迎你，<span style="color: #0e7a0d;"> {{user.userName}}</span>
                    </div>
                    <n-button strong secondary @click="tologin" style="position:absolute; right:40px; top:8px;">&#12288;退出登录&#12288;</n-button>
                <div class="space"></div>
                </div>
            </div>      
        </div>
    </div>
        <div class="main-panel">
            <div class="menus">

                <div v-if = "showSchRank" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><School></School></n-icon>&#8197;大学排名管理
                </div>
                <div v-else class="option" @click="toSchRank">
                    &#12288;&#12288;<n-icon size="20"><School></School></n-icon>&#8197;大学排名管理
                </div>
                <div class="space"></div>

                <div v-if = "showSubRank" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><Library></Library></n-icon>&#8197;学科排名管理
                </div>
                <div v-else class="option" @click="toSubRank">
                    &#12288;&#12288;<n-icon size="20"><Library></Library></n-icon>&#8197;学科排名管理
                </div>
                <div class="space"></div>

                <div v-if = "showSchool" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><Podium></Podium></n-icon>&#8197;大学信息管理
                </div>
                <div v-else class="option" @click="toSchool">
                    &#12288;&#12288;<n-icon size="20"><Podium></Podium></n-icon>&#8197;大学信息管理
                </div>
                <div class="space"></div>

                <div v-if = "showSubject" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><Book></Book></n-icon>&#8197;学科信息管理
                </div>
                <div v-else class="option" @click="toSubject">
                    &#12288;&#12288;<n-icon size="20"><Book></Book></n-icon>&#8197;学科信息管理
                </div>
                <div class="space"></div>

                <div v-if = "showPaper" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><Newspaper></Newspaper></n-icon>&#8197;论文数据管理
                </div>
                <div v-else class="option" @click="toPaper">
                    &#12288;&#12288;<n-icon size="20"><Newspaper></Newspaper></n-icon>&#8197;论文数据管理
                </div>
                <div class="space"></div>

                <div v-if = "showUser" class="myoption">
                    &#12288;&#12288;<n-icon size="20" color="#0e7a0d"><PersonCircle></PersonCircle></n-icon>&#8197;用户信息管理
                </div>
                <div v-else class="option" @click="toUser">
                    &#12288;&#12288;<n-icon size="20"><PersonCircle></PersonCircle></n-icon>&#8197;用户信息管理
                </div>
                <div class="space"></div>

            </div>
            <div style="padding:20px;width:100%">
                <router-view></router-view>
            </div>
        </div>
</template>

<script setup>
import {ref,reactive,inject, onMounted} from 'vue'
import {BarChart, School, Library, Podium, Book, Newspaper, PersonCircle} from "@vicons/ionicons5"

import {useRouter, useRoute} from 'vue-router'
const router = useRouter()
const route = useRoute()

const axios = inject("axios")
const message = inject("message")
const serverUrl = inject("serverUrl")
const dialog = inject("dialog")

const user = reactive({
    userName:""
})

const showSchRank = ref(false)
const showSubRank = ref(false)
const showSchool = ref(false)
const showSubject = ref(false)
const showPaper = ref(false)
const showUser = ref(false)

onMounted(() => {
    toSchRank()
    getInfo()
})

const getInfo = async() => {
    let res = await axios.get("/getInfo") 
    console.log(res)
    user.userName = res.data.data.user.UserName
}

const toSchRank = () => {
    showSchRank.value = true
    showSubRank.value = false
    showSchool.value = false
    showSubject.value = false
    showPaper.value = false
    showUser.value = false
    router.push("/dashboard/schrank")
}

const toSubRank = () => {
    showSchRank.value = false
    showSubRank.value = true
    showSchool.value = false
    showSubject.value = false
    showPaper.value = false
    showUser.value = false
    router.push("/dashboard/subrank")
}

const toSchool = () => {
    showSchRank.value = false
    showSubRank.value = false
    showSchool.value = true
    showSubject.value = false
    showPaper.value = false
    showUser.value = false
    router.push("/dashboard/school")
}

const toSubject = () => {
    showSchRank.value = false
    showSubRank.value = false
    showSchool.value = false
    showSubject.value = true
    showPaper.value = false
    showUser.value = false
    router.push("/dashboard/subject")
}

const toPaper = () => {
    showSchRank.value = false
    showSubRank.value = false
    showSchool.value = false
    showSubject.value = false
    showPaper.value = true
    showUser.value = false
    router.push("/dashboard/paper")
}

const toUser = () => {
    showSchRank.value = false
    showSubRank.value = false
    showSchool.value = false
    showSubject.value = false
    showPaper.value = false
    showUser.value = true
    router.push("/dashboard/user")
}

const tologin = () => {
    dialog.info({
        title: '退出登录',
        content: '是否要退出登录',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            router.push("/")
        },
        onNegativeClick: () => {}
    })    
}

</script>

<style lang="scss" scoped>
.container {
    .topbar {
        position: sticky;
        top: 0;
        height: 50px;
        background: white;
        border-bottom: 2px solid #efeff5;
    }
}
.main-panel {
    display: flex;
    color: #64676a;
    max-width: 1500px;
}
.menus {
    padding: 20px 10px;
    width: 250px;
    height: 80vh;
    border-right: 2px solid #efeff5;
    .myoption {
        padding: 10px 10px;
        box-sizing: border-box;
        line-height: 40px;
        text-align: left;
        border-radius: 5px;
        color: #0e7a0d;
        background-color: #e7f5ee;
    }
    .option {
        padding: 10px 10px;
        box-sizing: border-box;
        line-height: 40px;
        text-align: left;
        cursor: pointer;
        border-radius: 5px;
        &:hover {
            background-color: #efeff5;
        }
    }
    .space{
        height: 5px;
    }
}
</style>