<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="学科排名添加">
                <n-form ref="formRef" :rules="rules" :model="addSubRank" style="margin-top:10px;">
                    <n-form-item path="year" label="年份">
                        <n-input-number v-model:value="addSubRank.year" clearable  placeholder="请输入年份" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schName" label="大学名称">
                        <n-select v-model:value="addSubRank.schName" filterable :options="schoolOptions" clearable placeholder="请输入大学名称" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="subName" label="学科名称">
                        <n-select v-model:value="addSubRank.subName" filterable :options="subjectOptions" clearable placeholder="请输入学科名称" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="subCountRank" label="学科全国排名">
                        <n-input v-model:value="addSubRank.subCountRank" clearable placeholder="请输入学科全国排名" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="subWorldRank" label="学科全球排名">
                        <n-input v-model:value="addSubRank.subWorldRank" clearable placeholder="请输入学科全球排名" style="width: 500px;" />
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
            <n-tab-pane name="maintain" tab="学科排名维护">
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
                            <th>学科全国排名</th>
                            <th>学科全球排名</th>
                            <th>操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(subRank,index) in subRankList">
                            <td>{{subRank.ID}}</td>
                            <td>{{subRank.Year}}</td>
                            <td>{{subRank.SchName}}</td>
                            <td>{{subRank.SubName}}</td>
                            <td>{{subRank.SubCountRank}}</td>
                            <td>{{subRank.SubWorldRank}}</td>
                            <td>             
                                <n-button @click="updateValue(subRank)">
                                    <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                    修改&#8194;
                                </n-button>
                                <n-button @click="toDelete(subRank)" style="margin-left:20px;">
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
                <n-pagination @update:page="loadSubRank" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  

                <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form ref="formRef" :rules="rules" :model="updateSubRank">
                                <n-form-item label="年份" style="margin-top: 20px;">
                                    <n-input v-model:value="updateSubRank.year" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="大学名称" >
                                    <n-input v-model:value="updateSubRank.schName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="学科名称" >
                                    <n-input v-model:value="updateSubRank.subName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item path="subCountRank" label="学科全国排名">
                                    <n-input v-model:value="updateSubRank.subCountRank" clearable placeholder="请输入学科全国排名" />
                                </n-form-item>
                                <n-form-item path="subWorldRank" label="学科全球排名">
                                    <n-input v-model:value="updateSubRank.subWorldRank" clearable placeholder="请输入学科全球排名" />
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
const addSubRank = reactive({
    year : 2022,
    schName: null,
    subName: null,
    subCountRank: "",
    subWorldRank: ""
})
const updateSubRank = reactive({
    id: 0,
    year : null,
    schName: "",
    subName: "",
    subCountRank: "",
    subWorldRank: ""
})
const subRankList = ref([])
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
    loadSubRank()
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

const loadSubRank = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/subRank/search?year=${pageInfo.year}&schName=${pageInfo.schName}&subName=${pageInfo.subName}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        subRankList.value = res.data.data.subRank
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

function validateRank(rule, value) {
    if (showUpdateModel.value) {
        console.log(Number(value))
        console.log(updateSubRank.subCountRank)
        return Number(value) >= Number(updateSubRank.subCountRank)
    } else {
        console.log(Number(value))
        console.log(addSubRank.subCountRank)
        return Number(value) >= Number(addSubRank.subCountRank)
    }
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
        { required: true, message: "请输入学科名称" },
    ],
    subCountRank: [
        { required: true, message: "请输入学科全国排名" },
        { validator: onlyAllowNumber, message: "只能输入数字"},
    ],
    subWorldRank: [
        { required: true, message: "请输入学科全球排名" },
        { validator: onlyAllowNumber, message: "只能输入数字"},
        { validator: validateRank, message: "请输入正确排名"},
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
    console.log(addSubRank)
    let res = await axios.post("/subRank", {
        year: addSubRank.year,
        schName: addSubRank.schName,
        subName: addSubRank.subName,
        subCountRank: Number(addSubRank.subCountRank),
        subWorldRank: Number(addSubRank.subWorldRank)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadSubRank()  
    } else {
        message.error(res.data.msg)
    }
}

const updateValue = async(subRank) => {
    showUpdateModel.value = true
    updateSubRank.id = subRank.ID
    updateSubRank.year = subRank.Year
    updateSubRank.subName = subRank.SubName
    updateSubRank.schName = subRank.SchName
    updateSubRank.subCountRank = subRank.SubCountRank
    updateSubRank.subWorldRank = subRank.SubWorldRank
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
    formRef.value?.validate((errors) => {
        console.log(errors)
        if (errors) {
            message.error("数据格式有误")
            return
        }
    })
    let res = await axios.put("/subRank/"+ updateSubRank.id, {
        subCountRank: Number(updateSubRank.subCountRank),
        subWorldRank: Number(updateSubRank.subWorldRank)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadSubRank()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(subRank) => {
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/subRank/${subRank.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadSubRank()
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
    addSubRank.year = 2022
    addSubRank.subName = null
    addSubRank.schName = null
    addSubRank.subCountRank = ""
    addSubRank.subWorldRank = ""
}

const search = () => {
    pageInfo.pageNum = 1
    loadSubRank()
}

const clean = () => {
    pageInfo.year = null
    pageInfo.schName = ""
    pageInfo.subName = ""
    loadSubRank()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportSubRank?year=${pageInfo.year}&schName=${pageInfo.schName}&subName=${pageInfo.subName}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>
</style>