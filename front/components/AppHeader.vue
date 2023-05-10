<template>
  <v-app-bar app color="primary" dark height="40">
    <v-toolbar-title v-on:click="toHome">My App</v-toolbar-title>
    <v-spacer></v-spacer>
    <div class="mr-2">
      <v-menu>
        <template v-slot:activator="{ props }">
          <svg-icon
            type="mdi"
            :path="mdiAccount"
            v-bind="props"
            class="pointer"
          ></svg-icon>
        </template>
        <v-list>
          <v-list-item
            v-for="(item, index) in menuItems"
            :key="index"
            :value="index"
          >
            <v-list-item-title v-on:click="item.action">
              {{ item.title }}
            </v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </div>
  </v-app-bar>
</template>

<script setup lang="ts">
import SvgIcon from '@jamescoyle/vue-icon'
import { mdiAccount } from '@mdi/js'
import { useAuth } from '@/composables/useAuth'
const { currentUser, signout } = useAuth()

interface Menu {
  title: string
  action: () => void | Promise<void>
}

const router = useRouter()

const toHome = () => {
  router.push('/home')
}

const menuItems: Menu[] = [
  {
    title: currentUser.value.name ?? '',
    action: () => {
      console.log(currentUser.value.email)
    }
  },
  { title: 'sign out', action: signout }
]
</script>

<style scoped>
.pointer {
  cursor: pointer;
}
</style>
