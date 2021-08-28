<template>
  <v-md-preview :text="text" height="700px" style="text-align: left"></v-md-preview>
</template>

<script>
export default {
  name: 'Markdown',
  data() {
    return {
      title: '',
      author: '',
      digest: '',
      content: '',
      createdTime: ''
    };
  },
  async created(){

  }
};
</script>
