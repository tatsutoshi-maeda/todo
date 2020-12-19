import styled from 'styled-components'
import { List } from "./List"
import { Form } from "./Form"
import { useEffect, useState } from "react";
import { getLanguages } from "./const/languages";

const Header = styled.header`
  display: flex;
  justify-content: space-between;
  padding 24px 64px 0;
  border-bottom: 1px solid #E0E0E0;
`

const HeaderUl = styled.ul`
  display: flex;
  margin: 0;
  padding: 0;
`

const HeaderLi = styled.li`
  list-style: none;
  padding: 4px 12px;
  cursor: pointer;
  border-bottom: ${props => props.focused ? '2px solid #F44336' : 'none'};
`

function App() {

  const [tab, setTab] = useState('list');
  const [langs, setLangs] = useState([]);

  //functional ComponentだとuseEffectで制御する
  //useEffectはmounting時とupdating時に発火される
  useEffect(() => {
    console.log('App.js:useEffect');
    fetchLanguages();
  //下記のように配列を入れて何も入れなかった場合、mountingの時にのみ発火される
  },[])
  //下記のように、updateで依存するStateを決めることもできる
  // },[langs, tab])
  //下記のように何も入れなければ、mountingとupdatingで動く
  // })

  const fetchLanguages = async () => {
    const languages = await getLanguages();
    setLangs(languages);
  }

  const addLang = (lang) => {
    setLangs([...langs, lang]);
    setTab('list');
  }

  return (
    <div>
      <Header>
        <HeaderUl>
          <HeaderLi focused={tab === 'list'} onClick={() => setTab('list')}>リスト</HeaderLi>
          <HeaderLi focused={tab === 'form'} onClick={() => setTab('form')}>フォーム</HeaderLi>
          <HeaderLi focused={tab === 'book'} onClick={() => setTab('book')}>BOOK</HeaderLi>
        </HeaderUl>
      </Header>
      {
        // 三項演算子と即時関数以外でJSXに分岐書くやり方探す
        // tab === 'list' ? <List langs={langs}/> : <Form onAddLang={addLang}/>
        (() => {
          if (tab === 'list'){
            return <List langs={langs}/> 
          }else if (tab === 'form'){
            return <Form onAddLang={addLang}/>
          } else if (tab === 'book') {
            return <div>BOOK</div>
          }
        }) ()
      }
    </div>
  );
}

export default App;

//ReactのClass Componentでのライフサイクルの順番
/**
 *  Mounting(マウント時):Componentが表示されるまでの期間のこと。
  1.constructor()
  2.getDerivedStateFromProps() // 一般的ではないライフサイクル
  3.render() // ここで描画される
  4.componentDidMount()

 * Updating(更新時):Componentが表示され、Stateの更新を行う期間のこと。
  1.getDerivedStateFromProps() // 一般的ではないライフサイクル
  2.shouldComponentUpdate() // 一般的ではないライフサイクル
  3.render()
  4.getSnapshotBeforeUpdate() // 一般的ではないライフサイクル
  5.componentDidUpdate()

 * Unmounting(マウント解除時):他のComponentに移るときにだけ呼ばれる期間。
  1.componentWillUnmount()
 * */ 

 //呼び出し元になっているApp.jsはUnmountingされない
 //Unmountingのメソッドは遷移先の方に書く