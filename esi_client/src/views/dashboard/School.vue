<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="大学信息添加">
                <n-form ref="formRef" :rules="rules" :model="addSchool" style="margin-top:10px;">
                    <n-form-item path="id" label="大学编号">
                        <n-input v-model:value="addSchool.id" placeholder="请输入大学编号" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schName" label="大学名称">
                        <n-input v-model:value="addSchool.schName" placeholder="请输入大学名称" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item label="学校类别">
                        <n-select v-model:value="addSchool.schType" :options="options" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item label="办学性质">
                        <n-radio-group v-model:value="addSchool.schManage" name="radiogroup">
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
            <n-tab-pane name="maintain" tab="大学信息维护">
                <div style="display: flex; margin-top: 10px;">
                    <n-input v-model:value="pageInfo.schName" placeholder="请输入大学名称关键字" clearable />
                    <n-select v-model:value="pageInfo.schType" :options="options" style="width: 300px; margin-left: 20px;" placeholder="请选择学校类别" filterable />
                    <n-select v-model:value="pageInfo.schManage" :options="manage" style="width: 300px; margin-left: 20px;" placeholder="请选择办学性质" filterable />
                    <n-button @click="search" type="success" style="margin-left: 20px;">
                        <template #icon><n-icon><SearchOutline /></n-icon></template>
                        查询&#8194;
                    </n-button>
                </div>
                    <div style="display: flex; ">
                        <n-table :bordered="true" :single-line="false" style="margin-top:20px;">
                            <thead>
                                <tr>
                                <th>大学编号</th>
                                <th>大学名称</th>
                                <th>大学类别</th>
                                <th>办学性质</th>
                                <th>操作</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(school,index) in schoolList">
                                <td>{{school.ID}}</td>
                                <td>{{school.SchName}}</td>
                                <td>{{school.SchType}}</td>
                                <td>{{school.SchManage}}</td>
                                <td>             
                                    <n-button @click="updateValue(school)">
                                        <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                        修改&#8194;
                                    </n-button>
                                    <n-button @click="toDelete(school)" style="margin-left:20px;">
                                        <template #icon><n-icon><TrashOutline /></n-icon></template>                            
                                        删除&#8194;
                                    </n-button>
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
                        <n-pagination @update:page="loadSchool" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  
                        
                        <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form ref="formRef" :rules="rules" :model="updateSchool">
                                <n-form-item label="大学编号" style="margin-top: 20px;">
                                    <n-input v-model:value="updateSchool.id" :disabled="!active" />
                                </n-form-item>
                                <n-form-item path="schName" label="大学名称">
                                    <n-input v-model:value="updateSchool.schName" placeholder="请输入大学名称" clearable />
                                </n-form-item>
                                <n-form-item label="办学性质">
                                    <n-radio-group v-model:value="updateSchool.schManage" name="radiogroup">
                                        <n-space>
                                            <n-radio v-for="m in manage" :key="m.value" :value="m.value">
                                                {{ m.label }}&#12288;
                                            </n-radio>
                                        </n-space>
                                    </n-radio-group>
                                </n-form-item>
                                <n-form-item label="学校类别">
                                    <n-select v-model:value="updateSchool.schType" :options="options" />
                                </n-form-item>
                                <n-form-item label="" >
                                    <n-button @click="closeModal">取消</n-button>
                                    <n-button @click="toUpdate" type="success" style="margin-left: 20px;">确认</n-button>
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
const addSchool = reactive({
    id : "",
    schName: "",
    schType: "985/211",
    schManage: "公办大学"
})
const updateSchool = reactive({
    id : "",
    schName: "",
    schType: "",
    schManage: "",
})
const schoolList = ref([])
const pageInfo = reactive({
    pageNum:1,
    pageSize:7,
    pageCount:0,
    count:0,
    schName:"",
    schType:null,
    schManage:null
})

onMounted(() => {
    loadSchool()
})

const loadSchool = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/school/search?schName=${pageInfo.schName}&schType=${pageInfo.schType}&schManage=${pageInfo.schManage}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        schoolList.value = res.data.data.school
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

function onlyAllowNumber(rule, value) {
    return (/^\d*$/.test(value))
}

let rules = {
    id: [
        { required: true, message: "请输入大学编号" },
        { validator: onlyAllowNumber, message: "只能输入数字" },
    ],
    schName: [
        { required: true, message: "请输入大学名称" },
        { max: 20, message: "大学名称长度在小于 20 个字符" },    
    ]
}

const manage = ref([
    {value: "公办大学", label:"公办大学"},
    {value: "民办大学", label:"民办大学"},
    {value: "独立院校", label:"独立院校"},
])

const options = ref([
    {value:"985/211", label: "985/211"},
    {value:"211", label: "211"},
    {value:"中央部署本科院校", label: "中央部署本科院校"},
    {value:"省属本科院校", label: "省属本科院校"},
    {value:"高职（高专）院校", label: "高职（高专）院校"},
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
    let res = await axios.post("/school", {
        id: Number(addSchool.id),
        schName: addSchool.schName,
        schType: addSchool.schType,
        schManage: addSchool.schManage
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadSchool()  
    } else {
        message.error(res.data.msg)
    }
}

const updateValue = async(school) => {
    showUpdateModel.value = true
    updateSchool.id = school.ID
    updateSchool.schName = school.SchName
    updateSchool.schType = school.SchType
    updateSchool.schManage = school.SchManage
}

const toUpdate = async() => {
    formRef.value?.validate((errors) => {
        console.log(errors)
        if (errors) {
            message.error("数据格式有误")
        } else {
            update();
        }
    })
}

const update = async() => {
    let res = await axios.put("/school/"+ updateSchool.id, {
        schName: updateSchool.schName,
        schType: updateSchool.schType,
        schManage: updateSchool.schManage
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadSchool()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(school) => {
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/school/${school.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadSchool()
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
    addSchool.id = ""
    addSchool.schName = ""
    addSchool.schType = "985/211"
    addSchool.schManage = "公办大学"
}

const search = () => {
    pageInfo.pageNum = 1
    loadSchool()
}

const clean = () => {
    pageInfo.schName = ""
    pageInfo.schType = null
    pageInfo.schManage = null
    pageInfo.pageNum = 1
    loadSchool()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportSchool?schName=${pageInfo.schName}&schType=${pageInfo.schType}&schManage=${pageInfo.schManage}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>

</style>