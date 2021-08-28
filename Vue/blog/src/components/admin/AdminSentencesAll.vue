<template>
  <suspense>
    <template #default>
      <AdminSentencesAllPage />
    </template>
    <template #fallback>
      加载页面中....
    </template>
  </suspense>
</template>

<script>
import AdminSentencesAllPage from "./AdminSentencesAllPage";

export default {
  name: "AdminSentencesAll",
  components: {AdminSentencesAllPage}
}
</script>

<style scoped>

</style>
