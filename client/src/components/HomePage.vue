<template>
  <div class="home">
    <h1>{{ msg }}</h1>

    <div class="max-w-md bg-white rounded-lg shadow-lg p-8">
      <img
        alt="Fateskink logo"
        src="./../assets/logo.png"
        class="centered-image"
        style="width: 150px; height: 150px"
      />
      <h1 class="text-3xl font-bold text-gray-800 mb-2">Name</h1>
      <p class="text-gray-600 mb-4">Description about yourself</p>
      <div class="flex justify-center">
        <a
          href="#"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mr-2"
          >Button 1</a
        >
        <a
          href="#"
          class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >Button 2</a
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HomePage",
  props: {
    msg: String,
  },
};
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
