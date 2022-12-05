<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="论文数据添加">
                <n-form ref="formRef" :rules="rules" :model="addPaper" style="margin-top:10px;">
                    <n-form-item path="year" label="年份">
                        <n-input-number v-model:value="addPaper.year" clearable  placeholder="请输入年份" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schName" label="大学名称">
                        <n-select v-model:value="addPaper.schName" filterable :options="schoolOptions" clearable placeholder="请输入大学名称" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="subName" label="学科名称">
                        <n-select v-model:value="addPaper.subName" filterable :options="subjectOptions" clearable placeholder="请输入学科名称" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="paperNum" label="论文数">
                        <n-input v-model:value="addPaper.paperNum" clearable placeholder="请输入论文数" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="usedNum" label="他引数">
                        <n-input v-model:value="addPaper.usedNum" clearable placeholder="请输入他引数" style="width: 500px;" />
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
            <n-tab-pane name="maintain" tab="论文数据维护">
                <div style="display: flex; margin-top: 10px;">
                    <n-input-number v-model:value="pageInfo.year" clearable  placeholder="请输入年份" style="width: 500px;" />
                    <n-input v-model:value="pageInfo.schName" placeholder="请输入大学名称关键字" clearable style="margin-left: 20px;" />
                    <n-input v-model:value="pageInfo.subName" placeholder="请输入学科名称关键字" clearable style="margin-left: 20px;" />
                    <n-button @click="search" type="success" style="margin-left: 20px;">
                        <template #icon><n-icon><SearchOutline /></n-icon></template>
                            查询&#8194;
                    </n-button>
                </div>
                <div style="display: flex; ">
                    <n-table :bordered="true" :single-line="false" style="margin-top:20px;">
                        <thead>
                            <tr>
                            <th>编号</th>
                            <th>年份</th>
                            <th>大学名称</th>
                            <th>学科名称</th>
                            <th>论文数</th>
                            <th>他引数</th>
                            <th>操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(paper,index) in paperList">
                            <td>{{paper.ID}}</td>
                            <td>{{paper.Year}}</td>
                            <td>{{paper.SchName}}</td>
                            <td>{{paper.SubName}}</td>
                            <td>{{paper.PaperNum}}</td>
                            <td>{{paper.UsedNum}}</td>
                            <td>             
                                <n-button @click="toUpdate(paper)">
                                    <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                    修改&#8194;
                                </n-button>
                                <n-button @click="toDelete(paper)" style="margin-left:20px;">
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
                <n-pagination @update:page="loadPaper" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  

                <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form ref="formRef" :rules="rules" :model="updatePaper">
                                <n-form-item label="年份" style="margin-top: 20px;">
                                    <n-input v-model:value="updatePaper.year" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="大学名称" >
                                    <n-input v-model:value="updatePaper.schName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="学科名称" >
                                    <n-input v-model:value="updatePaper.subName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item path="paperNum" label="论文数">
                                    <n-input v-model:value="updatePaper.paperNum" clearable placeholder="请输入论文数" />
                                </n-form-item>
                                <n-form-item path="UsedNum" label="他引数">
                                    <n-input v-model:value="updatePaper.usedNum" clearable placeholder="请输入他引数" />
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
const addPaper = reactive({
    year : 2022,
    schName: null,
    subName: null,
    paperNum: "",
    usedNum: ""
})
const updatePaper = reactive({
    id: 0,
    year : null,
    schName: "",
    subName: "",
    paperNum: "",
    usedNum: ""
})
const paperList = ref([])
const pageInfo = reactive({
    pageNum:1,
    pageSize:7,
    pageCount:0,
    count:0,
    year:null,
    schName: "",
    subName: ""
})
const schoolOptions = ref([])
const subjectOptions = ref([])

onMounted(() => {
    loadSchoolList()
    loadSubjectList()
    loadPaper()
})

const loadSchoolList = async() => {
    let res = await axios.get("school")
    console.log(res)
    schoolOptions.value = res.data.data.school.map((item)=>{
        return {
            label: item.SchName,
            value: item.SchName,
        }
    })
}

const loadSubjectList = async() => {
    let res = await axios.get("subject")
    console.log(res)
    subjectOptions.value = res.data.data.subject.map((item)=>{
        return {
            label: item.SubName,
            value: item.SubName,
        }
    })
}

const loadPaper = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/paper/search?year=${pageInfo.year}&schName=${pageInfo.schName}&subName=${pageInfo.subName}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        paperList.value = res.data.data.paper
    }
    pageInfo.count = res.data.data.count;
    pageInfo.pageCount = parseInt(pageInfo.count / pageInfo.pageSize) + (pageInfo.count % pageInfo.pageSize > 0 ? 1 : 0)
}

function validateYear(rule, value) {
    let now = new Date()
    let year = now.getFullYear()
    return value <= year
}

function onlyAllowNumber(rule, value) {
    return (/^\d*$/.test(value))
}

let rules = {
    year: [
        { required: true, message: "请输入年份" },
        { validator: validateYear, message: "请输入正确年份"},
    ],
    schName: [
        { required: true, message: "请输入大学名称" },
    ],
    subName: [
        { required: true, message: "请输入学科名称"},
    ],
    paperNum: [
        { required: true, message: "请输入论文数" },
        { validator: onlyAllowNumber, message: "只能输入数字" },
    ],
    usedNum: [
        { required: true, message: "请输入他引数" },
        { validator: onlyAllowNumber, message: "只能输入数字" },
    ],
}

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
    console.log(addPaper)
    let res = await axios.post("/paper", {
        year: addPaper.year,
        schName: addPaper.schName,
        subName: addPaper.subName,
        paperNum: Number(addPaper.paperNum),
        usedNum: Number(addPaper.usedNum)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadPaper()  
    } else {
        message.error(res.data.msg)
    }
}

const toUpdate = async(paper) => {
    showUpdateModel.value = true
    updatePaper.id = paper.ID
    updatePaper.year = paper.Year
    updatePaper.subName = paper.SubName
    updatePaper.schName = paper.SchName
    updatePaper.paperNum = paper.PaperNum
    updatePaper.usedNum = paper.UsedNum
}

const update = async() => {
    let res = await axios.put("/paper/"+ updatePaper.id, {
        paperNum: Number(updatePaper.paperNum),
        usedNum: Number(updatePaper.usedNum)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadPaper()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(paper) => {
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/paper/${paper.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadPaper()
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
    addPaper.year = 2022
    addPaper.subName = null
    addPaper.schName = null
    addPaper.paperNum = ""
    addPaper.usedNum = ""
}

const search = () => {
    pageInfo.pageNum = 1
    loadPaper()
}

const clean = () => {
    pageInfo.year = null
    pageInfo.schName = ""
    pageInfo.subName = ""
    loadPaper()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportPaper?year=${pageInfo.year}&schName=${pageInfo.schName}&subName=${pageInfo.subName}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>
</style>