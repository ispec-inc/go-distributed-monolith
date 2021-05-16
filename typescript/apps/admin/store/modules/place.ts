import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule,
} from 'vuex-module-decorators'
import { GetRequest } from '@monorepo/gen/admin/api/rest/article/article_pb'
import { $axios } from '~/util/api'
import { store } from '~/store'
import PlaceData from '~/types/place'

@Module({
  name: 'modules/place',
  dynamic: true,
  store,
  namespaced: true,
})
export class PlaceModule extends VuexModule {
  list: Array<PlaceData> = []
  current: PlaceData | null = null
  request = new GetRequest()

  public get placeList(): PlaceData[] | null {
    return this.list ?? null
  }

  public get currentPlace(): PlaceData | null {
    return this.current ?? null
  }

  @Mutation
  private SET_LIST(value: PlaceData[]) {
    this.list = value
  }

  @Mutation
  private SET_CURRENT(value: PlaceData) {
    this.current = value
  }

  @Action({ rawError: true })
  public fetchPlaceList(): Promise<void> {
    return $axios
      .get<PlaceData[]>('/places')
      .then((response) => {
        console.log(response)
        // this.SET_LIST(response.data)
      })
      .catch((e) => {
        console.log(e)
      })
  }

  @Action({ rawError: true })
  public fetchPlace(): Promise<void> {
    return $axios
      .get<PlaceData>('/places')
      .then((response) => {
        console.log(response)
        // this.SET_CURRENT(response.data)
      })
      .catch((e) => {
        console.log(e)
      })
  }
}

const placeModule = getModule(PlaceModule)
export default placeModule
