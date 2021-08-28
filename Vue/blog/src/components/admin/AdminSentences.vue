<template>
  <suspense>
    <template #default>
      <AdminSentencesPage />
    </template>
    <template #fallback>
      加载页面中....
    </template>
  </suspense>
</template>

<script>

import Sentences from "../../views/Sentences";
import AdminSentencesPage from "./AdminSentencesPage";
export default {
  name: "AdminSentences",
  components: {AdminSentencesPage, Sentences},
}
</script>

<style scoped>

</style>
