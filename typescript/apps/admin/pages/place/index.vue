<template>
  <div>
    <v-card>
      <a-data-table
        :title="'プレイス一覧'"
        :headers="headers"
        :items="places"
        :loading="isFetching"
        @detail="goDetailPage"
        @edit="goEditPage"
        @delete="openDialogDelete"
      ></a-data-table>
    </v-card>
    <dialog-delete
      :dialog="dialog"
      :target="target"
      @cancel="dialog = false"
      @delete="deleteItem"
    ></dialog-delete>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import PlaceResource from '~/resources/place'
import Resource from '~/resources/resource'
import placeModule from '~/store/modules/place'
import PlaceData from '~/types/place'

@Component({})
export default class PlacePage extends Vue {
  isFetching = false
  dialog = false
  target: Resource | null = null
  headers = [
    {
      text: 'ID',
      value: 'id',
    },
    {
      text: '名前(日本語)',
      value: 'nameJa',
    },
    {
      text: 'タイトル(英語)',
      value: 'nameEn',
    },
    {
      text: '編集/削除',
      value: 'actions',
      align: 'center',
      sortable: false,
    },
  ]

  get places(): PlaceData[] {
    return placeModule?.placeList ?? []
  }

  async mounted() {
    this.isFetching = true
    await placeModule.fetchPlaceList()
    this.isFetching = false
  }

  deleteItem() {
    if (this.target) {
      this.dialog = false
    }
  }

  goDetailPage(place: PlaceResource) {
    this.$router.push(`/place/${place.id}`)
  }

  goEditPage(place: PlaceResource) {
    this.$router.push(`/place/${place.id}/edit`)
  }

  openDialogDelete(place: PlaceResource) {
    this.target = place
    this.dialog = true
  }
}
</script>
