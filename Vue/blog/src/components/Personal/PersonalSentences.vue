<template>
  <header> 全部文章</header>
  <el-divider style="margin: 4px"></el-divider>

  <suspense>
    <template #default>
      <PersonalSentencesPage />
    </template>
    <template #fallback>
      加载中...
    </template>
  </suspense>
</template>

<script>
import PersonalSentencesPage from "./PersonalSentencesPage";

export default {
  name: "PersonalSentences",
  components:{
    PersonalSentencesPage
  }
}
</script>

<style scoped>

</style>
