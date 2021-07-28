import React, { useState, useEffect } from 'react';
import { createStyles, fade, makeStyles, Theme } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { Content, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    table: {
      minWidth: 650,
    },
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.black, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.black, 0.25),
      },
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: 'auto',
      },
    },
    searchIcon: {
      padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    inputRoot: {
      color: 'inherit',
    },
    inputInput: {
      padding: theme.spacing(1, 1, 1, 0),
      // vertical padding + font size from searchIcon
      paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: '30ch',
        '&:focus': {
          width: '20ch',
        },
      },
    },
  }),
);

export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [users, setUser] = React.useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    const getUser = async () => {
      const res = await http.listUser({ limit: 20, offset: 0 });
      setLoading(false);
      setUser(res);
    };
    getUser();
  }, [loading]);




  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={'แสดงข้อมูลพนักงาน'}
        subtitle="แสดงข้อมูลพนักงาน">
        <Button color="secondary" variant="contained" href="/">BACK</Button>
      </Header>

      <Content>

        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">รหัสพนักงาน</TableCell>
                <TableCell align="center">ชื่อ-นามสกุล</TableCell>
                <TableCell align="center">เพศ</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {users.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.idemployee}</TableCell>
                  <TableCell align="center">{item.name}</TableCell>
                  <TableCell align="center">{item.edges?.gender?.gender}</TableCell>
                  <TableCell align="center">
                    
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>

  );
}