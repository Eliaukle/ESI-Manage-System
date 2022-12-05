<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="用户信息添加">
                <n-form ref="formRef" :rules="rules" :model="addUser" style="margin-top:10px;">
                    <n-form-item path="userName" label="用户名">
                        <n-input v-model:value="addUser.userName" placeholder="请输入用户名" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="password" label="密码">
                        <n-input v-model:value="addUser.password" type="password" placeholder="请输入密码" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="repeatPassword" label="确认密码">
                        <n-input v-model:value="addUser.repeatPassword" type="password" placeholder="请再次输入密码" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item label="用户身份">
                        <n-radio-group v-model:value="addUser.identify" name="radiogroup">
                            <n-space>
                                <n-radio v-for="m in manage" :key="m.value" :value="m.value">
                                    {{ m.label }}&#12288;
                                </n-radio>
                            </n-space>
                        </n-radio-group>
                    </n-form-item>
                    <n-form-item label="">
                        <n-button @click="toCreate">
                            <template #icon><n-icon><AddCircleOutline /></n-icon></template>
                            添加&#8194;
                        </n-button>
                        <n-button @click="refresh" style="margin-left: 30px;">
                            <template #icon><n-icon><RefreshCircleOutline /></n-icon></template>                            
                            重置
                        </n-button>
                    </n-form-item>
                </n-form>               
            </n-tab-pane>
            <n-tab-pane name="maintain" tab="用户信息维护">
                <div style="display: flex; margin-top: 10px;">
                    <n-input v-model:value="pageInfo.userName" placeholder="请输入用户名关键字" clearable />
                    <n-select v-model:value="pageInfo.identify" :options="manage" style="width: 300px; margin-left: 20px;" placeholder="请选择用户身份" filterable />
                    <n-button @click="search" type="success" style="margin-left: 20px;">
                        <template #icon><n-icon><SearchOutline /></n-icon></template>
                        查询&#8194;
                    </n-button>
                </div>
                    <div style="display: flex; ">
                        <n-table :bordered="true" :single-line="false" style="margin-top:20px;">
                            <thead>
                                <tr>
                                <th>用户编号</th>
                                <th>用户名</th>
                                <th>用户身份</th>
                                <th>操作</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(user,index) in userList">
                                <td>{{user.ID}}</td>
                                <td>{{user.UserName}}</td>
                                <td>{{user.Identify}}</td>
                                <td>
                                    <div v-if="curUser.id != user.ID">
                                        <n-button @click="toUpdate(user)">
                                            <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                            修改&#8194;
                                        </n-button> 
                                        <n-button @click="toDelete(user)" style="margin-left:20px;">
                                            <template #icon><n-icon><TrashOutline /></n-icon></template>                            
                                            删除&#8194;
                                        </n-button>
                                    </div>  
                                    <div v-else>
                                        <n-button disabled >
                                            <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                            修改&#8194;
                                        </n-button> 
                                        <n-button style="margin-left:20px;" disabled >
                                            <template #icon><n-icon><TrashOutline /></n-icon></template>                            
                                            删除&#8194;
                                        </n-button>
                                    </div>
                                </td>
                                </tr>
                            </tbody>
                        </n-table>
                        <n-button @click="clean" style="margin-top:20px;margin-left: 20px;">
                            <template #icon><n-icon><RefreshCircleOutline /></n-icon></template>
                            重置&#8194;
                        </n-button>
                        <n-button @click="exportExcel" style="margin-top:20px;margin-left: 20px;">
                            <template #icon><n-icon><DownloadOutline /></n-icon></template>
                            导出&#8194;
                        </n-button>
                        </div>
                        <n-pagination @update:page="loadUser" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  
                        <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form :rules="rules" :model="updateUser">
                                <n-form-item label="用户名" style="margin-top: 20px;">
                                    <n-input v-model:value="updateUser.userName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="用户身份">
                                    <n-radio-group v-model:value="updateUser.identify" name="radiogroup">
                                        <n-space>
                                            <n-radio v-for="m in manage" :key="m.value" :value="m.value">
                                                {{ m.label }}&#12288;
                                            </n-radio>
                                        </n-space>
                                    </n-radio-group>
                                </n-form-item>
                                <n-form-item label="" >
                                    <n-button @click="closeModal">取消</n-button>
                                    <n-button @click="update" type="success" style="margin-left: 20px;">确认</n-button>
                                </n-form-item>
                            </n-form>   
                        </n-modal>

            </n-tab-pane>
        </n-tabs>
    </div>
</template>

<script setup>
import {ref,reactive,inject, onMounted} from 'vue'
import {AddCircleOutline, RefreshCircleOutline, SearchOutline, TrashOutline, CreateOutline, DownloadOutline} from "@vicons/ionicons5"

import {useRouter, useRoute} from 'vue-router'
const router = useRouter()
const route = useRoute()

const axios = inject("axios")
const message = inject("message")
const serverUrl = inject("serverUrl")
const dialog = inject("dialog")
const fileDownload = inject("fileDownload")

const formRef = ref(null)
const showUpdateModel = ref(false)
const curUser = reactive({
    id: "",
    identify: "",
})
const addUser = reactive({
    userName: "",
    password: "",
    repeatPassword:"",
    identify: "系统管理员",
})
const updateUser = reactive({
    id: "",
    userName: "",
    identify: "",
})
const userList = ref([])
const pageInfo = reactive({
    pageNum:1,
    pageSize:7,
    pageCount:0,
    count:0,
    userName:"",
    identify:null
})

onMounted(() => {
    getInfo()
    loadUser()
})

const getInfo = async() => {
    let res = await axios.get("/getInfo") 
    console.log(res)
    curUser.id = res.data.data.user.ID
    curUser.identify = res.data.data.user.Identify 
}

const loadUser = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/user/search?userName=${pageInfo.userName}&identify=${pageInfo.identify}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        userList.value = res.data.data.user
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

function validatePasswordSame(rule, value) {
    return value == addUser.password;
}

let rules = {
    userName: [
        { required: true, message: "请输入用户名" },
        { min: 3, max: 20, message: "用户名长度在 3 到 20 个字符" },
    ],
    password: [
        { required: true, message: "请输入密码" },
        { min: 6, max: 20, message: "密码长度在 6 到 20 个字符" },    
    ],
    repeatPassword: [
        { required: true, message: "请重新输入密码" }, 
        { validator: validatePasswordSame, message: "两次输入的密码不一致" },
    ],
}

const manage = ref([
    {value: "系统管理员", label:"系统管理员"},
    {value: "普通用户", label:"普通用户"},
])

const toCreate = async() => {
    formRef.value?.validate((errors) => {
        if (errors) {
            message.error("数据格式有误")
        } else {
            create();
        }
    })
}

const create = async() => {
    if (curUser.identify == "普通用户" && addUser.identify == "系统管理员") {
        message.error("权限不足")
        return
    }
    let res = await axios.post("/user", {
        userName: addUser.userName,
        password: addUser.password,
        identify: addUser.identify
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadUser()  
    } else {
        message.error(res.data.msg)
    }
}

const toUpdate = async(user) => {
    if (curUser.identify == "普通用户" && user.Identify == "系统管理员") {
        message.error("权限不足")
        return
    }
    showUpdateModel.value = true
    updateUser.id = user.ID
    updateUser.userName = user.UserName
    updateUser.identify = user.Identify
}

const update = async() => {
    if (curUser.identify == "普通用户" && updateUser.identify == "系统管理员") {
        message.error("权限不足")
        return
    }
    let res = await axios.put("/user/"+ updateUser.id, {
        identify: updateUser.identify
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadUser()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(user) => {
    if (curUser.identify == "普通用户" && user.Identify == "系统管理员") {
        message.error("权限不足")
        return
    }
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/user/${user.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadUser()
            }else{
                message.error(res.data.msg)
            }  
        },
        onNegativeClick: () => {}
    })    
}

const closeModal = () => {
    showUpdateModel.value = false;
}

const refresh = () => {
    addUser.userName = ""
    addUser.password = ""
    addUser.repeatPassword = ""
    addUser.identify = "系统管理员"
}

const search = () => {
    pageInfo.pageNum = 1
    loadUser()
}

const clean = () => {
    pageInfo.userName = ""
    pageInfo.identify = null
    pageInfo.pageNum = 1
    loadUser()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportUser?userName=${pageInfo.userName}&identify=${pageInfo.identify}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>

</style>