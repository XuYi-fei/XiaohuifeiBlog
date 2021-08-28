<template>
  <suspense>
    <template #default>
      <AdminBlogsAllPage />
    </template>
    <template #fallback>
      加载中...
    </template>
  </suspense>
</template>

<script>
import AdminBlogsAllPage from "./AdminBlogsAllPage";
export default {
  name: "AdminBlogsAll",
  components: {AdminBlogsAllPage}
}
</script>

<style scoped>

</style>
