import React, {useEffect, useState} from 'react';
import {
    BrowserRouter as Router,
    Route,
    Switch,
    useHistory
} from "react-router-dom";
import 'react-bulma-components/dist/react-bulma-components.min.css';
import {Section, Button} from "react-bulma-components";
import axios from 'axios';
import {Pie} from 'react-chartjs-2';

const Home = ({filter, setFilter}) => {
    let history = useHistory();
    return (
        <Section className={"section"}>
            <div className={"columns is-centered is-mobile"}>
                <div className={"column is-one-quarter-desktop is-half-tablet"}>
                    <div className={"box"}>
                        <div className={"field"}>
                            <h1>Filter (empty for all)</h1>
                        </div>
                        <div className={"field"}>
                            <div className={"control"}>
                                <input className={"input is-primary"} value={filter}
                                       onChange={setFilter}/>
                            </div>
                        </div>
                        <Button className={"button is-primary"} onClick={() => history.push("/search")}>
                            Search
                        </Button>
                    </div>
                </div>
            </div>
        </Section>
    );
};

function getRandomColor() {
    let letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}

const Search = ({filter}) => {
    const [repos, setRepos] = useState([]);
    const [loaded, setLoaded] = useState(false);
    const [pie, setPie] = useState();

    useEffect(() => {
        const fetchData = async () => {
            const result = await axios.post(
                'http://localhost:8765/search', {filter}
            );
            setRepos(result.data);
            let lgs = [];
            let lines = [];
            let colors = [];
            if (result.data) {
                result.data.forEach((repo) => {
                    if (repo.Language) {
                        repo.Language.forEach((l, idx) => {
                            if (lgs.includes(l)) {
                                lines[lgs.indexOf(l)] += repo.Lines[idx];
                            } else {
                                lgs.push(l);
                                lines.push(repo.Lines[idx]);
                                colors.push(getRandomColor());
                            }
                        });
                    }
                });
                setPie({
                    labels: lgs,
                    datasets: [{
                        data: lines,
                        backgroundColor: colors
                    }]
                });
            }
            setLoaded(true);
        };
        fetchData();
    }, [filter]);

    let i = 0;

    if (loaded) {
        return (
            <div>
                <div className={"field"}>
                    {filter !== "" ?
                        <h1 className={"title"}>
                            Repository matching "<b>{filter}</b>"
                        </h1> :
                        <h1 className={"title"}>
                            Last 100 repositories
                        </h1>
                    }
                </div>
                {repos ?
                    <div className={"columns"}>
                        <div className={"column"}>
                            <div className={"box"}>
                                <h1>Lines per language</h1>
                                <Pie data={pie} height={500}/>
                            </div>
                        </div>
                        <div className={"column"}>
                            {repos.map(repo =>
                                <div className={"box"} key={i++}>
                                    <div className={"columns"}>
                                        <div className={"column"}>
                                            <b>
                                                {repo.Name}
                                            </b>
                                        </div>
                                        <div className={"column"}>
                                            <a href={repo.URL}>
                                                {repo.URL}
                                            </a>
                                        </div>
                                    </div>
                                    <div>
                                        {repo.Language ?
                                            <div className={"columns is-multiline"}>
                                                {repo.Language.map((l, idx) =>
                                                    <div key={idx} className={"column"}>
                                                        <div className={"box"}>
                                                            <b>{l}<br/></b>
                                                            <i>{repo.Lines[idx]} lines</i>
                                                        </div>
                                                    </div>
                                                )}
                                            </div> : <div>Undefined</div>
                                        }
                                    </div>
                                </div>
                            )}
                        </div>
                    </div> : <div>No matches found</div>}
            </div>);
    } else {
        return (
            <div>
                <h1>Loading repositories</h1>
            </div>
        );
    }
};

function App() {
    const [filter, setFilter] = useState('');

    const handleFilter = event => setFilter(event.target.value);

    return (
        <Router>
            <Switch>
                <Route exact path="/"><Home filter={filter} setFilter={handleFilter}/></Route>
                <Route path="/search"><Search filter={filter}/></Route>
            </Switch>
        </Router>
    );
}

export default App;
