<template>
  <div id="field-div">
      <ul id="nav-field-ul">
        <li class="classification-item">
          <el-button plain="plain"  size="medium" autofocus="autofocus" name="0" style="width: 160px; margin: 1px" @click="newPage('recentBlogs', 'latest')">最新更新 </el-button>
        </li>
        <li class="classification-item"  v-for="(item, index) in fields">
          <el-button plain="plain" type="primary" name="{{index}}" size="medium" autofocus="autofocus" style="width: 160px; margin: 1px" @click="changesStatus(index)"> {{ item.fieldName}} </el-button>
          <ul id="nav-classification-ul" v-show="item.show">
            <li class="classification-item" v-for="classItem in classification">
              <el-badge :value="classItem.num" class="item" type="primary">
                <el-button plain="plain" autofocus="autofocus" size="small" style="width: 120px; margin: 1px" @click="newPage(item.fieldName, classItem.classificationName)">{{ classItem.classificationName}}</el-button>
              </el-badge>
            </li>
          </ul>
        </li>
      </ul>
  </div>

</template>

<script>
import {defineComponent, reactive, ref} from "vue";
import {getClassification, getField} from "../api/classification";

export default defineComponent({
  name: "ClassificationPage",
  async setup(){
    const response = await getField(  "all", "")
    const fields = reactive(response.data)
    const classification = reactive([])
    const plain = ref(true)
    const autofocus = ref(true)
    return{
      autofocus: autofocus,
      plain: plain,
      fields: fields,
      classification: classification
    }
  },
  methods:{
    async changesStatus(i){
      const response = await getClassification("all", "", this.fields[i].fieldName)
      for (let k = 0; k < this.fields.length; k++){
        if(i === k)
          this.fields[k].show = (this.fields[k].show === undefined? true: !this.fields[k].show)
        else
          this.fields[k].show = false
      }
      this.classification = response.data
    },
    newPage(field, classification){
      window.location = '/classification/' + field + '/' + classification
    }
  }
})
</script>


<style scoped>
div#field-div{
  display: flex;
  flex-direction: row-reverse;
  min-width: 145px;
  margin: 10px;
}

ul{
  list-style: none;
}

ul#nav-field-ul{
  position: sticky;
  top: 50px;
  width: 162px;
  display: flex;
  flex-direction: column;
  margin: 0;
  padding: 0;
}

li.classification-item{
  display: inline;
}

span.classification-item{
  font-size: 17px;
  padding: 3px;
}
</style>
