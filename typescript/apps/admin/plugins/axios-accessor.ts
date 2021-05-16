import { Plugin } from '@nuxt/types'
import { initializeAxios } from '~/util/api'

export const accessor: Plugin = ({ $axios }): void => {
  initializeAxios($axios)
}

export default accessor
