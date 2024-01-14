// import logo from './logo.svg';
// import { useState, useEffect } from 'react';
// import { useState } from 'react';
import './App.css';
import Header from './components/CommonPages/header.jsx';
// import PostList from './components/Pages/ClientPages/PostList.jsx';
// import { ShowPost } from './components/Pages/ClientPages/showPost.jsx';
// import { PostShow } from './components/Pages/ClientPages/PostShow.jsx';
// import { BrowserRouter as router,Routes,Route,link } from 'react-router-dom';
// import Login from './components/Pages/GuestPages/Login.jsx';
// import SignUp from './components/Pages/GuestPages/SignUp.jsx';
import NewPostList from './components/Pages/ClientPages/NewPostList.jsx';
import SamplePage from './components/Pages/SamplePage/SamplePage.jsx';


function App() {
  // const [cookies, setCookies] = useState(0)
  // const [count, setCount] = useState(0)

  // useEffect(() => { // runs only on time when compnent mounted
  //   const id = setInterval(() => setCount((oldCount) => oldCount + 1), 1000);

  //   return () => {
  //     clearInterval(id);
  //   };
  // }, []);

  // useEffect(() => {
  //   console.log('Hey cookies got changed', cookies)
  // }, [cookies]);
  // const [selectedPost, setSelectedPost] = useState(null);

  // const onPostSelect = (post) => {
  //   console.log(post)
  //   setSelectedPost(post);
  // }

  return (

    <div className="App">

      {true && <Header />}

      {/* <PostList onSelectPost={(post) => { console.log(post); setSelectedPost(post) }} />
      <ShowPost {...selectedPost} /> */}

      {/* 
      <PostList onPostSelect={onPostSelect} />
      {selectedPost && <PostShow {...selectedPost} />} */}
      {/* {count}
      {cookies && <p>Hey im cookies {cookies}</p>}
      <button onClick={() => { setCookies(!cookies) }}>Click me</button> */}
      <SamplePage />
      <NewPostList />

    </div>
  );
}

export default App;
