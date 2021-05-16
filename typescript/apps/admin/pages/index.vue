<template>
  <div>
    <div>トップ</div>
    <!--    {{ placeList }}-->
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import articleModule from '@/store/modules/article.ts'

@Component({})
export default class TopPage extends Vue {
  created() {
    articleModule.fetchArticles()
  }

  // get placeList() {
  //   return placeModule.placeList
  // }
}
</script>
