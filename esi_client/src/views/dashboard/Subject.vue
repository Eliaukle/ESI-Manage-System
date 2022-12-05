<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="学科信息添加">
                <n-form ref="formRef" :rules="rules" :model="addSubject" style="margin-top:10px;">
                    <n-form-item path="subName" label="学科名称">
                        <n-input v-model:value="addSubject.subName" placeholder="请输入学科名称" clearable style="width: 500px;" />
                    </n-form-item>
                    <n-form-item label="学科类别">
                        <n-select v-model:value="addSubject.subType" :options="options" style="width: 500px;" />
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
            <n-tab-pane name="maintain" tab="学科信息维护">
                <div style="display: flex; margin-top: 10px;">                  
                    <n-input v-model:value="pageInfo.subName" placeholder="请输入学科名称关键字" clearable />
                    <n-select v-model:value="pageInfo.subType" :options="options" style="width: 300px; margin-left: 20px;" placeholder="请选择学科类别" filterable />
                    <n-button @click="search" type="success" style="margin-left: 20px;">
                        <template #icon><n-icon><SearchOutline /></n-icon></template>
                        查询&#8194;
                    </n-button>
                </div>
                    <div style="display: flex; ">
                        <n-table :bordered="true" :single-line="false" style="margin-top:20px;">
                            <thead>
                                <tr>
                                <th>学科编号</th>
                                <th>学科名称</th>
                                <th>学科类别</th>
                                <th>操作</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(subject,index) in subjectList">
                                <td>{{subject.ID}}</td>
                                <td>{{subject.SubName}}</td>
                                <td>{{subject.SubType}}</td>
                                <td>             
                                    <n-button @click="updateValue(subject)">
                                        <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                        修改&#8194;
                                    </n-button>
                                    <n-button @click="toDelete(subject)" style="margin-left:20px;">
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
                        <n-pagination @update:page="loadSubject" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  
                        
                        <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form ref="formRef" :rules="rules" :model="updateSubject">
                                <n-form-item label="学科编号" style="margin-top: 20px;">
                                    <n-input v-model:value="updateSubject.id" :disabled="!active" />
                                </n-form-item>
                                <n-form-item path="subName" label="学科名称">
                                    <n-input v-model:value="updateSubject.subName" placeholder="请输入大学名称" clearable />
                                </n-form-item>
                                <n-form-item label="学科类别">
                                    <n-select v-model:value="updateSubject.subType" :options="options" />
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
const addSubject = reactive({
    subName: "",
    subType: "哲学"
})
const updateSubject = reactive({
    id: 0,
    subName: "",
    subType: ""
})
const subjectList = ref([])
const pageInfo = reactive({
    pageNum:1,
    pageSize:7,
    pageCount:0,
    count:0,
    subName:"",
    subType:null
})

onMounted(() => {
    loadSubject()
})

const loadSubject = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/subject/search?subName=${pageInfo.subName}&subType=${pageInfo.subType}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        subjectList.value = res.data.data.subject
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

let rules = {
    subName: [
        { required: true, message: "请输入学科名称" },
        { max: 20, message: "学科名称长度在小于 20 个字符" },    
    ]
}

const options = ref([
    {value:"哲学", label: "哲学"},
    {value:"经济学", label: "经济学"},
    {value:"法学", label: "法学"},
    {value:"教育学", label: "教育学"},
    {value:"文学", label: "文学"},
    {value:"历史学", label: "历史学"},
    {value:"理学", label: "理学"},
    {value:"工学", label: "工学"},
    {value:"农学", label: "农学"},
    {value:"医学", label: "医学"},
    {value:"军事学", label: "军事学"},
    {value:"管理学", label: "管理学"},
    {value:"艺术学", label: "艺术学"},
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
    let res = await axios.post("/subject", {
        subName: addSubject.subName,
        subType: addSubject.subType,
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadSubject()  
    } else {
        message.error(res.data.msg)
    }
}

const updateValue = async(subject) => {
    showUpdateModel.value = true
    updateSubject.id = subject.ID
    updateSubject.subName = subject.SubName
    updateSubject.subType = subject.SubType
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
    let res = await axios.put("/subject/"+ updateSubject.id, {
        subName: updateSubject.subName,
        subType: updateSubject.subType,
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadSubject()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(subject) => {
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/subject/${subject.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadSubject()
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
    addSubject.subName = ""
    addSubject.subType = "哲学"
}

const search = () => {
    pageInfo.pageNum = 1
    loadSubject()
}

const clean = () => {
    pageInfo.subName = ""
    pageInfo.subType = null
    loadSubject()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportSubject?subName=${pageInfo.subName}&subType=${pageInfo.subType}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>

</style>