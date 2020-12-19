import { List } from "./List"
import { Form } from "./Form"
import { useEffect, useState } from "react";
import { getLanguages } from "./const/languages";

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

  let style = {
    backgroundColor: "gray",
    color: "#FFF",
    margin: 10,
    padding: 10,
}

  return (
    <div style={style}>
      <header>
        <ul>
          <li onClick={() => setTab('list')}>リスト</li>
          <li onClick={() => setTab('form')}>フォーム</li>
        </ul>
      </header>
      <hr/>
      {
        tab === 'list' ? <List langs={langs}/> : <Form onAddLang={addLang}/>
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