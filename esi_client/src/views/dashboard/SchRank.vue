<template>
    <div>
        <n-tabs type="line" animated>
            <n-tab-pane name="add" tab="大学排名添加">
                <n-form ref="formRef" :rules="rules" :model="addSchRank" style="margin-top:10px;">
                    <n-form-item path="year" label="年份">
                        <n-input-number v-model:value="addSchRank.year" clearable  placeholder="请输入年份" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schName" label="大学名称">
                        <n-select v-model:value="addSchRank.schName" filterable :options="options" clearable placeholder="请输入大学名称" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schCountRank" label="大学全国排名">
                        <n-input v-model:value="addSchRank.schCountRank" clearable placeholder="请输入大学全国排名" style="width: 500px;" />
                    </n-form-item>
                    <n-form-item path="schWorldRank" label="大学全球排名">
                        <n-input v-model:value="addSchRank.schWorldRank" clearable placeholder="请输入大学全球排名" style="width: 500px;" />
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
            <n-tab-pane name="maintain" tab="大学排名维护">
                <div style="display: flex; margin-top: 10px;">
                    <n-input-number v-model:value="pageInfo.year" clearable  placeholder="请输入年份" style="width: 500px;" />
                    <n-input v-model:value="pageInfo.schName" placeholder="请输入大学名称关键字" clearable style="margin-left: 20px;" />
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
                            <th>大学全国排名</th>
                            <th>大学全球排名</th>
                            <th>操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(schRank,index) in schRankList">
                            <td>{{schRank.ID}}</td>
                            <td>{{schRank.Year}}</td>
                            <td>{{schRank.SchName}}</td>
                            <td>{{schRank.SchCountRank}}</td>
                            <td>{{schRank.SchWorldRank}}</td>
                            <td>             
                                <n-button @click="updateValue(schRank)">
                                    <template #icon><n-icon><CreateOutline /></n-icon></template>                            
                                    修改&#8194;
                                </n-button>
                                <n-button @click="toDelete(schRank)" style="margin-left:20px;">
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
                <n-pagination @update:page="loadSchRank" v-model:page="pageInfo.pageNum" :page-count="pageInfo.pageCount" style="margin-top: 20px;"/>  

                <n-modal v-model:show="showUpdateModel" preset="dialog" title="Dialog">
                            <template #header>
                                <div>修改</div>
                            </template>
                            <n-form ref="formRef" :rules="rules" :model="updateSchRank">
                                <n-form-item label="年份" style="margin-top: 20px;">
                                    <n-input v-model:value="updateSchRank.year" :disabled="!active" />
                                </n-form-item>
                                <n-form-item label="大学名称" >
                                    <n-input v-model:value="updateSchRank.schName" :disabled="!active" />
                                </n-form-item>
                                <n-form-item path="schCountRank" label="大学全国排名">
                                    <n-input v-model:value="updateSchRank.schCountRank" clearable placeholder="请输入大学全国排名" />
                                </n-form-item>
                                <n-form-item path="schWorldRank" label="大学全球排名">
                                    <n-input v-model:value="updateSchRank.schWorldRank" clearable placeholder="请输入大学全球排名" />
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
const addSchRank = reactive({
    year : 2022,
    schName: null,
    schCountRank: "",
    schWorldRank: ""
})
const updateSchRank = reactive({
    id: 0,
    year : null,
    schName: "",
    schCountRank: "",
    schWorldRank: ""
})
const schRankList = ref([])
const pageInfo = reactive({
    pageNum:1,
    pageSize:7,
    pageCount:0,
    count:0,
    year:null,
    schName: ""
})
const options = ref([])

onMounted(() => {
    loadSchoolList()
    loadSchRank()
})

const loadSchoolList = async() => {
    let res = await axios.get("school")
    console.log(res)
    options.value = res.data.data.school.map((item)=>{
        return {
            label: item.SchName,
            value: item.SchName,
        }
    })
}

const loadSchRank = async(pageNum = 0) => {
    if (pageNum != 0){
        pageInfo.pageNum = pageNum;
    }
    let res = await axios.post(`/schRank/search?year=${pageInfo.year}&schName=${pageInfo.schName}&pageNum=${pageInfo.pageNum}&pageSize=${pageInfo.pageSize}`)
    console.log(res)
    if (res.data.code == 200) {
        schRankList.value = res.data.data.schRank
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
        return Number(value) >= Number(updateSchRank.schCountRank)
    } else {
        return Number(value) >= Number(addSchRank.schCountRank)
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
    schCountRank: [
        { required: true, message: "请输入大学全国排名" },
        { validator: onlyAllowNumber, message: "只能输入数字" },
    ],
    schWorldRank: [
        { required: true, message: "请输入大学全球排名" },
        { validator: onlyAllowNumber, message: "只能输入数字" },
        { validator: validateRank, message: "请输入正确排名" },
    ],
}

const toCreate = async() => {
    formRef.value?.validate((errors) => {
        console.log(errors)
        if (errors) {
            message.error("数据格式有误")
        } else {
            create();
        }
    })
}

const create = async() => {
    console.log(addSchRank)
    let res = await axios.post("/schRank", {
        year: addSchRank.year,
        schName: addSchRank.schName,
        schCountRank: Number(addSchRank.schCountRank),
        schWorldRank: Number(addSchRank.schWorldRank)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        refresh()
        loadSchRank()  
    } else {
        message.error(res.data.msg)
    }
}

const updateValue = async(schRank) => {
    showUpdateModel.value = true
    updateSchRank.id = schRank.ID
    updateSchRank.year = schRank.Year
    updateSchRank.schName = schRank.SchName
    updateSchRank.schCountRank = schRank.SchCountRank
    updateSchRank.schWorldRank = schRank.SchWorldRank
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
    let res = await axios.put("/schRank/"+ updateSchRank.id, {
        schCountRank: Number(updateSchRank.schCountRank),
        schWorldRank: Number(updateSchRank.schWorldRank)
    })
    console.log(res)
    if (res.data.code == 200) {
        message.success(res.data.msg)
        loadSchRank()  
        closeModal()
    } else {
        message.error(res.data.msg)
    }    
}

const toDelete = async(schRank) => {
    dialog.warning({
        title: '删除',
        content: '是否要删除',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
            let res = await axios.delete(`/schRank/${schRank.ID}`)
            if(res.data.code == 200){
                message.success(res.data.msg)
                loadSchRank()
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
    addSchRank.year = 2022
    addSchRank.schName = null
    addSchRank.schCountRank = ""
    addSchRank.schWorldRank = ""
}

const search = () => {
    pageInfo.pageNum = 1
    loadSchRank()
}

const clean = () => {
    pageInfo.year = null
    pageInfo.schName = ""
    loadSchRank()
}

const exportExcel = async() => {
    let res = await axios.get(`/exportSchRank?year=${pageInfo.year}&schName=${pageInfo.schName}`, {responseType: "blob"})
    console.log(res)
    fileDownload(res.data, "result.xlsx")
}

</script>

<style lang="scss" scoped>
</style>