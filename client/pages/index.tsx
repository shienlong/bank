import Link from 'next/link';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Grid from '@material-ui/core/Grid';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
  }),
);


const Home = () => {
  const classes = useStyles({});

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            BANK
          </Typography>
        </Toolbar>
      </AppBar>
      <Grid item xs={12} md={6}>
        <Grid container spacing={1} direction="column" alignItems="center">
          <Grid item>
            <ButtonGroup variant="contained" color="secondary" size="large" >
              <Button>
                <Link href="/balance" as={`/balance`}>Balance</Link>
              </Button>
              <Button>
                <Link href="/transfer" as={`/transfer`}>Transfer</Link>
              </Button>
              <Button>
                <Link href="/invest" as={`/invest`}>Invest</Link>
              </Button>
            </ButtonGroup>
          </Grid>
        </Grid>
      </Grid>
    </div>
  )
}

export default Home
