<script lang="ts" setup>
import {GetBreedList, GetImageUrlByBreed, GetRandomImageUrl} from "../../wailsjs/go/main/App";
import {onMounted, reactive, Ref, ref} from "vue";

let randomImageUrl = ref("")
let breeds: Ref<string[]> = ref([]);
let photos: Ref<string[]> = ref([]);
let selectedBreed: Ref<string> = ref("");
let showRandomPhoto: Ref<boolean> = ref(false)
let showBreedPhotos: Ref<boolean> = ref(false)



const getRandomImageUrl = async () => {
  showRandomPhoto.value = false;
  showBreedPhotos.value = false;
  randomImageUrl.value = await GetRandomImageUrl();
  showRandomPhoto.value = true;
}

const getBreedList =  async () => {
  const breedsList = await GetBreedList();
  breeds.value = breedsList
}

onMounted(() => {
  getBreedList();
})


const getImageUrlByBreed = async () => {
  showRandomPhoto.value = false;
  showBreedPhotos.value = false;
  console.log(selectedBreed.value)
  photos.value = await GetImageUrlByBreed(selectedBreed.value);
  showBreedPhotos.value = true
}

</script>

<template>
  <main>
    <h3>Dog API</h3>
    <button class="btn" @click="getRandomImageUrl()">Fetch a dog randomly</button>
    Click on down arrow to select a breed
    <select v-model="selectedBreed">
      <option v-for="breed in breeds" :key="breed" :value="breed">{{breed}}</option>
    </select>
    <button class="btn" @click="getImageUrlByBreed()">Fetch photos by breed</button>
    <br>
    <div v-if="showRandomPhoto">
      <img :src="randomImageUrl" alt="No dog found">
    </div>
    <div v-if="showBreedPhotos">
      <div v-for="photo in photos" :key="photo">
        <img :src="photo" alt="No dog found">
      </div>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
