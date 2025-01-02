import axios from 'axios'


export const search = async (query: string) => {
  try {
    const url = `http://localhost:3333/search?query=${query}`

    const { data }  = await axios.get(url)

    return data

  } catch (error) {
    return []
  }
}
