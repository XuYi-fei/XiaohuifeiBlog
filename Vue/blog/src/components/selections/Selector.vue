<template>
  <el-select v-model="value" placeholder="请选择">
    <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
    </el-option>
  </el-select>
</template>

<script>
export default {
  name: "Selector",
  data() {
    return {
      options: [{
        value: 'time',
        label: '最近修改'
      }, {
        value: 'degree',
        label: '最热博文'
      }],
      value: 'time'
    }
  }
}
</script>

<style scoped>

</style>
