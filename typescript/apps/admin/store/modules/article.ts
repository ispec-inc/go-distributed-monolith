import {
  Module,
  VuexModule,
  Mutation,
  Action,
  getModule,
} from 'vuex-module-decorators'
import { GetRequest } from '@monorepo/gen/admin/api/rest/article/article_pb'
import { $axios } from '@/util/api'
import { store } from '@/store'
import PlaceData from '@/types/place'

@Module({
  name: 'modules/article',
  dynamic: true,
  store,
  namespaced: true,
})
export class ArticleModule extends VuexModule {
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
  public fetchArticles(): Promise<void> {
    return $axios
      .get<PlaceData[]>('/v1/articles')
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

const articleModule = getModule(ArticleModule)
export default articleModule
