import { useState, useContext, useEffect } from 'react'
import { API_URL } from '../../constants'
import { useRouter } from 'next/router'
import { AuthContext, UserInfo } from '../../modules/auth_provider'

const index = () => {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [isSignup, setIsSignup] = useState(false)
  const { authenticated } = useContext(AuthContext)



  const router = useRouter()

  useEffect(() => {
    if (authenticated) {
      router.push('/')
      return
    }
  }, [authenticated])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()


    if (isSignup && password !== confirmPassword) {
      alert('Password and Confirm Password do not match')
      return
    }

    try {
      const endpoint = isSignup ? '/auth/signup' : '/auth/login'
      const body = isSignup ? { username, email, password, confirmPassword} : { email, password }
      const res = await fetch(`${API_URL}${endpoint}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
      })

      const data = await res.json()
      if (res.ok) {
        const user: UserInfo = {
          username: data.username,
          id: data.id,
        }

        localStorage.setItem('user_info', JSON.stringify(user))
        return router.push('/')
      }
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div className='flex items-center justify-center min-w-full min-h-screen'>
      <form className='flex flex-col md:w-1/5' onSubmit={submitHandler}>
        <div className='text-3xl font-bold text-center'>
          <span className='text-blue'>{isSignup ? 'Sign Up' : 'Welcome!'}</span>
        </div>
        <input
          placeholder='email'
          className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        {isSignup && (
          <>
            <input
              placeholder='username'
              className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </>
        )}
        <input
          type='password'
          placeholder='password'
          className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        {isSignup && (
            <>
              <input
                type='password'
                placeholder='confirm password'
                className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
              />
            </>
        )}

        <button
          className='p-3 mt-6 rounded-md bg-blue font-bold text-white'
          type='submit'
        >
          {isSignup ? 'Sign Up' : 'Login'}
        </button>
        <button
          className='p-3 mt-2 rounded-md bg-gray-200 font-bold text-black'
          type='button'
          onClick={() => setIsSignup(!isSignup)}
        >
          {isSignup ? 'Already have an account? Login' : 'Don\'t have an account? Sign Up'}
        </button>
      </form>
    </div>
  )
}


export default index