import axios from 'axios'


export const info = async () => {
  try {
    const url = `http://localhost:3333/info`

    const { data }  = await axios.get(url)

    return data

  } catch (error) {
    return []
  }
}
