<template>
  <el-empty description="暂无数据" style="height: 980px" v-show="empty"></el-empty>
  <ul>
    <li v-for="item in sentences">
      <div id="block">
        <div id="arrangement">
          <p class="sentence-content" style="white-space: pre-line">{{item.content}}</p>
          <p class="sentence-author">{{ item.author === ''? ' . ' : '—— '+item.author}}</p>
          <p class="sentence-user">{{item.userName}}</p>

        </div>
        <div id="operation">
            <span>
              <el-popconfirm title="确定删除吗？" @confirm="remove(item)">
                <template #reference>
                <el-button>删除</el-button>
                </template>
              </el-popconfirm>
              <el-popconfirm title="确定发布吗？" @confirm="pass(item)">
                <template #reference>
                <el-button>发布</el-button>
                </template>
              </el-popconfirm>
            </span>
        </div>

      </div>
      <el-divider style="margin: 5px"></el-divider>
    </li>
  </ul>
  <el-pagination
      @size-change="handleSizeChange"
      @current-change="getPageSentences"
      v-model:currentPage="currentPage"
      layout="prev, pager, next, jumper"
      :page-count="pageNum"
      :key="pageNum"
  >
  </el-pagination>
</template>

<script>
import {defineComponent, reactive, ref} from "vue";
import {getAllSentences, getSentencesPageNum} from "../../api/sentences";
import {DeleteSentencesTmp, PassSentences} from "../../api/admin";

export default defineComponent({
  name: "AdminSentencesPage",
  async setup(){
    let empty = ref(false)
    let pageNum = ref(1)
    let currentPage = ref(1)
    let sentences = reactive([])
    let response = await getSentencesPageNum("allTmp", "")
    pageNum = response.data.pageNum
    response = await getAllSentences(1, "allTmp", "")
    if(response.status !== 0){
      empty = true
    }else{
      sentences = response.data
    }
    return {
      empty: empty,
      pageNum: pageNum,
      currentPage: currentPage,
      sentences: sentences
    }
  },
  methods: {
    async pass(item){
      const response = await PassSentences(item)
      if (response.status !== 0){
        alert(response.msg)
      }else{
        alert(response.msg)
        window.location.reload()
      }
    },
    async remove(item){
      console.log(item)
      const response = await DeleteSentencesTmp(item)
      if (response.status !== 0){
        alert(response.msg)
      }else{
        alert(response.msg)
        window.location.reload()
      }
    },
    async getPageSentences(val){
      const response = await getAllSentences(val, "allTmp", "")
      this.sentences = response.data
      // this.refreshBlog()
      this.$forceUpdate()
    },
    handleSizeChange(){

    }
  }
})
</script>

<style scoped>
div#block{
  display: flex;
  justify-content: space-around;
}

div#arrangement{
  flex: 80%;
}

div#operation{
  flex: 20%;
  text-align: right;
  display: flex;
  flex-direction: column-reverse;
}

ul{
  list-style: none;
  text-align: left;
  flex: 80%;
}

p.sentence-content{
  font-size: large;
}

p.sentence-author{
  text-align: right;
}
</style>
